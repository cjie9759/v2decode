package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
	// _ "net/http/pprof"
)

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
	for _, v := range links {
		p := v
		go down(p, wg)
	}
	wg.Wait()
}

func down(u string, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	res, err := http.Get(u)
	if err == nil {
		defer res.Body.Close()

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
			}
		}

		if err != nil {
			fmt.Println(err, u)
			return
		}
		{
			// p :=
			ur, _ := url.Parse(u)

			fmt.Printf("https://%s  %.1f M/s\n", ur.Host, 100/time.Since(start).Seconds())
		}
	} else {
		fmt.Println(u, err, "is down")
	}
}
