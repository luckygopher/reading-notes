// @Description: 1.5 获取url内容
// @Author: Arvin
// @date: 2021/1/27 9:05 下午
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	urls := os.Args[1:]
	for _, url := range urls {
		data, respStatus, respStatusCode := fetch1(url)
		fmt.Printf("fetch %s data:%s respStatus:%s respStatusCode:%d", url, string(data), respStatus, respStatusCode)
	}
}

// ioutil.ReadAll
func fetch1(url string) ([]byte, string, int) {
	if !strings.HasPrefix(url, "http://") {
		if urls := strings.Split(url, "//"); len(urls) > 1 {
			url = urls[1]
		}
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch %s err:%v", url, err)
		// 终止进程，并返回一个status错误码
		os.Exit(1)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	respStatus := resp.Status
	respStatusCode := resp.StatusCode
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch %s read body err:%v", url, err)
		os.Exit(1)
	}
	return b, respStatus, respStatusCode
}

// io.Copy
func fetch2(url string) {
	if !strings.HasPrefix(url, "http://") {
		if urls := strings.Split(url, "//"); len(urls) > 1 {
			url = urls[1]
		}
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch %s err:%v", url, err)
		os.Exit(1)
	}
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		fmt.Fprintf(os.Stderr, "fetch %s copy body stdout err:%v", url, err)
		os.Exit(1)
	}
}
