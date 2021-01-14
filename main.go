package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	var url string
	var thread int
	flag.StringVar(&url, "url", "", "attack url")
	flag.IntVar(&thread, "t", 1, "thread nums")
	flag.Parse()
	fmt.Println(url)
	fmt.Println(thread)
	if len(url) == 0 {
		return
	}
	fmt.Println("start:", thread)

	for i := 0; i <= thread; i++ {
		go func(idx int) {
			for {
				start := time.Now().UnixNano()
				resp, _ := http.Get(url)
				bytes, _ := ioutil.ReadAll(resp.Body)
				end := time.Now().UnixNano()
				fmt.Println(idx, "reads :", len(bytes), "speed :", float64((end-start)/1e6))
				resp.Body.Close()
			}
		}(i)
	}
	var v int
	fmt.Scan(&v)
}
