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

}
func decode() {

	ReadFile(func(s *bufio.Scanner) {
		// if !strings.HasPrefix(s,`vmess://`) {
		// 	return
		// }
		t, err := base64.StdEncoding.DecodeString(s.Text()[8:])
		if err != nil {
			fmt.Println(s.Text())
			return
		}
		fmt.Println(string(t))
	})

}
func ReadFile(fun func(*bufio.Scanner)) {
	f, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	b := bytes.NewBuffer(f)

	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		// fmt.Printf("%s\n", scanner.Text())
		fun(scanner)
	}

}
