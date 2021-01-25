// @Description: 1.2 命令行参数
// @Author: Arvin
// @date: 2021/1/25 11:26 上午
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	echo1()
	echo2()
	echo3()
	echo4()
}

func echo1() {
	now := time.Now()
	var s, sep string
	// 每次循环连接追加，当参数数量庞大时，这种方式代价高昂
	for i := 1; i < len(os.Args); i++ {
		// 赋值运算符
		s += sep + os.Args[i]
		// 等价于 s = s + sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	sub := time.Now().Sub(now)
	fmt.Println("我是+=的耗时", sub)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	now := time.Now()
	// 使用 strings 包的 join 函数拼接字符串
	fmt.Println(strings.Join(os.Args[1:], " "))
	sub := time.Now().Sub(now)
	fmt.Println("我是join的耗时", sub)
}

func echo4() {
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
}
