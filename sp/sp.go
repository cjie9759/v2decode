package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
	// _ "net/http/pprof"
)

type Sp struct {
	Uri   string
	Speed float64
	Host  string
	Filsh bool
}

func main() {
	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()
	links :=
		strings.Split(`https://jnb-za-ping.vultr.com/vultr.com.100MB.bin,https://blr-in-ping.vultr.com/vultr.com.100MB.bin,https://del-in-ping.vultr.com/vultr.com.100MB.bin,https://bom-in-ping.vultr.com/vultr.com.100MB.bin,https://osk-jp-ping.vultr.com/vultr.com.100MB.bin,https://sel-kor-ping.vultr.com/vultr.com.100MB.bin,https://sgp-ping.vultr.com/vultr.com.100MB.bin,https://tlv-il-ping.vultr.com/vultr.com.100MB.bin,https://hnd-jp-ping.vultr.com/vultr.com.100MB.bin,https://mel-au-ping.vultr.com/vultr.com.100MB.bin,https://syd-au-ping.vultr.com/vultr.com.100MB.bin,https://ams-nl-ping.vultr.com/vultr.com.100MB.bin,https://fra-de-ping.vultr.com/vultr.com.100MB.bin,https://lon-gb-ping.vultr.com/vultr.com.100MB.bin,https://mad-es-ping.vultr.com/vultr.com.100MB.bin,https://man-uk-ping.vultr.com/vultr.com.100MB.bin,https://par-fr-ping.vultr.com/vultr.com.100MB.bin,https://sto-se-ping.vultr.com/vultr.com.100MB.bin,https://waw-pl-ping.vultr.com/vultr.com.100MB.bin,https://ga-us-ping.vultr.com/vultr.com.100MB.bin,https://il-us-ping.vultr.com/vultr.com.100MB.bin,https://tx-us-ping.vultr.com/vultr.com.100MB.bin,https://hon-hi-us-ping.vultr.com/vultr.com.100MB.bin,https://lax-ca-us-ping.vultr.com/vultr.com.100MB.bin,https://mex-mx-ping.vultr.com/vultr.com.100MB.bin,https://fl-us-ping.vultr.com/vultr.com.100MB.bin,https://nj-us-ping.vultr.com/vultr.com.100MB.bin,https://wa-us-ping.vultr.com/vultr.com.100MB.bin,https://sjo-ca-us-ping.vultr.com/vultr.com.100MB.bin,https://tor-ca-ping.vultr.com/vultr.com.100MB.bin,https://scl-cl-ping.vultr.com/vultr.com.100MB.bin,https://sao-br-ping.vultr.com/vultr.com.100MB.bin`, ",")

	blackhole, err := os.CreateTemp("", "blackhole")
	if err != nil {
		panic(err)
	}
	defer blackhole.Close()

	wg := &sync.WaitGroup{}
	wg.Add(len(links))
	sps := make([]*Sp, len(links))
	for k := range links {
		s := &Sp{Uri: links[k]}
		sps[k] = s
		go s.down(wg)
	}

	go func() {
		for {
			fmt.Print("\033[H\033[2J")
			sort.Slice(sps, func(i, j int) bool {
				return sps[i].Speed > sps[j].Speed
			})
			for k := range sps {
				if sps[k].Filsh {
					fmt.Printf("%s %.1f MB/s --------------------\n", sps[k].Host, sps[k].Speed)
				} else {
					fmt.Printf("%s		%.1f MB/s\n", sps[k].Host, sps[k].Speed)
				}
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()
	wg.Wait()
}

func (s *Sp) down(wg *sync.WaitGroup) {
	defer wg.Done()

	{
		ur, _ := url.Parse(s.Uri)
		s.Host = `https://` + ur.Host
	}

	start := time.Now()
	res, err := http.Get(s.Uri)
	if err == nil {
		defer res.Body.Close()
		i := 0
		{
			b := make([]byte, 1024)
			r := bufio.NewReader(res.Body)
			for {
				_, err := r.Read(b)
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatalf("Error reading file: %v", err)
					break
				}
				i++
				s.Speed = float64(i>>10) / time.Since(start).Seconds()
				// fmt.Printf("%.1f MB/s\n", s.Speed)
			}
		}

		if err != nil {
			fmt.Println(err, s.Uri)
			return
		}
		s.Filsh = true
		// fmt.Printf("https://%s  %.1f M/s\n", s.Host, 100/time.Since(start).Seconds())
	} else {
		fmt.Println(s.Uri, err, "is down")
	}
}
