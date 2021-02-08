// @Description: 1.7 web服务
// @Author: Arvin
// @date: 2021/1/28 4:55 下午
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/reading-notes/go-bible/ch1/lissajous"
)

func main() {
	//server1()
	//server2()
	//server3()
	server4()
}

func server1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var mu sync.Mutex
var count int

func server2() {
	// 返回请求路径
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count++
		mu.Unlock()
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})
	// 返回调用请求次数
	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		fmt.Fprintf(w, "Count %d\n", count)
		mu.Unlock()
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func server3() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
		fmt.Fprintf(w, "Host = %q\n", r.Host)
		fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func server4() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cycles, err := strconv.ParseFloat(r.FormValue("cycles"), 64)
		if err != nil {
			fmt.Errorf("参数错误:%v", err)
		}
		lissajous.Lissajous(w, cycles)
	})
	http.ListenAndServe("localhost:8000", nil)
}
