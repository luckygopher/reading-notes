package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

// 1byte = 8bit
// 计算两个SHA256哈希码中不同bit的数目
func countBit(str1, str2 string) int {
	var num int
	s1 := sha256.Sum256([]byte(str1))
	s2 := sha256.Sum256([]byte(str2))
	for i, b := range s1 {
		for j := 0; j < 8; j++ { // 对比bit位是否相同
			if (b & (1 << j)) != (s2[i] & (1 << j)) {
				num++
			}
		}
	}
	return num
}

// 打印标准输入，支持命令行flag定制。输出SHA256和SHA512编码
func InSHA256() {
	var str string
	flag.StringVar(&str, "str", "", "需要加密字符串")
	flag.Parse()
	if str == "" {
		stdin := bufio.NewScanner(os.Stdin)
		if stdin.Scan() {
			str = stdin.Text()
		}
	}
	fmt.Println("str:" + str)
	sha256 := sha256.Sum256([]byte(str))
	sha512 := sha512.Sum512([]byte(str))
	fmt.Printf("sha256:%x\nsha512:%x", sha256, sha512)
}

// 切片比较
func equal(x, y string) bool {
	if len(x) != len(y) {
		return false
	}
	for i, _ := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

// 回文数
func cover(a int64) int64 {
	var c int64
	var i int64 = 1
	for a > 0 {
		b := a % 10
		c = c*10 + b*i
		a = a / 10
	}
	return c
}
