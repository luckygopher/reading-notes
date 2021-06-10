// @Description:
// @Author: Arvin
// @date: 2021/5/31 2:59 下午
package main

import (
	"fmt"

	"golang.org/x/net/html"
)

func main() {
	// doc, err := html.Parse(os.Stdin)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
	// 	os.Exit(1)
	// }
	// for _, link := range visit(nil, doc) {
	// 	fmt.Println(link)
	// }
	fmt.Println(panicAndRecover(5))
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func panicAndRecover(n int) (result int, err error) {
	defer func() {
		switch r := recover(); r {
		case "gt10":
			err = fmt.Errorf("panic error %v", r)
		case nil:
			result = 10
		default:
			panic(r)
		}
	}()
	if n > 10 {
		panic("gt10")
	} else if n < 4 {
		panic("lt4")
	}
	return
}
