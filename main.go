package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func init() {
	// 获取命令行参数
	// fmt.Println("命令行参数数量:", len(os.Args))
	// for k, v := range os.Args {
	// 	fmt.Printf("args[%v]=[%v]\n", k, v)
	// }

	if len(os.Args) == 1 {
		fmt.Println("请输入文件名")
		os.Exit(0)
	}
	if len(os.Args) == 2 {
		path = os.Args[1]
		return
	}
	for _, v := range os.Args {
		if !strings.HasPrefix(v, `-`) {
			path = v
		}
	}

}

var path string

func main() {
	if len(os.Args) == 2 {
		decode()
		return
	}
	encode()

}
func decode() {
	var s *bufio.Scanner
	{
		f, err := os.ReadFile(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		b1, err := base64.StdEncoding.DecodeString(string(f))
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		b := bytes.NewBuffer(b1)
		s = bufio.NewScanner(b)
	}
	for s.Scan() {
		// fmt.Printf("%s\n", s.Text())
		if strings.HasPrefix(s.Text(), `vmess://`) {
			t, err := base64.StdEncoding.DecodeString(s.Text()[8:])
			if err != nil {
				fmt.Println(s.Text())
				return
			}
			fmt.Println(string(t))
		} else {
			fmt.Println(s.Text())
		}
	}
}
func encode() {
	f, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	b := bytes.NewBuffer(f)
	res := bytes.NewBuffer(nil)
	s := bufio.NewScanner(b)
	for s.Scan() {
		// fmt.Printf("%s\n", scanner.Text())
		if !strings.HasPrefix(s.Text(), `{`) {
			_, err := res.Write(s.Bytes())
			if err != nil {
				fmt.Println(err)
			}
		} else {
			res.WriteString(`vmess://`)
			res.Write([]byte(base64.StdEncoding.EncodeToString(s.Bytes())))
		}
		res.WriteString("\n")
	}
	// fmt.Println(res.String())
	fmt.Println(base64.StdEncoding.EncodeToString(res.Bytes()[:res.Len()-len("\n")]))
}
