// @Description: 1.6 并发获取多个url
// @Author: Arvin
// @date: 2021/1/28 10:49 上午
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type fetchResp struct {
	respStatus string
	byteNum    int64
}

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	//selfMain()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

// 自己实现
func selfMain() {
	fetchResult := make(chan fetchResp)
	urls := os.Args[1:]
	for _, url := range urls {
		go func(url string) {
			fetchResult <- selfFetch(url)
		}(url)
	}
	for range urls {
		fmt.Printf("%+v\n", <-fetchResult)
	}
}

func selfFetch(url string) fetchResp {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch %s err:%v", url, err)
		os.Exit(1)
	}
	byteNum, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "copy %s err:%v", url, err)
		os.Exit(1)
	}
	respStatus := resp.Status
	resp.Body.Close()
	return fetchResp{respStatus: respStatus, byteNum: byteNum}
}
