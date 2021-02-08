// @Description: 1.3 查找重复的行
// @Author: Arvin
// @date: 2021/1/25 11:26 上午
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//dup1()
	dup2()
	//dup3()
}

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	// 统计重复行
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors form input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("行：%s\t重复次数：%d\n", line, n)
		}
	}
}

func dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("行：%s\t重复次数：%d\n", line, n)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] != 0 {
			fmt.Printf("文件:%s\t出现重复行\t内容:%s\n", f.Name(), input.Text())
		}
		counts[input.Text()]++
	}
}

func dup3() {
	counts := make(map[string]int)
	files := os.Args[1:]
	// 追加标准输入，如果命令行参数为空则从标准输入获取
	if len(files) == 0 {
		files = append(files, "/dev/stdin")
	}
	for _, file := range files {
		byteData, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read file err:%v", err)
			continue
		}
		data := strings.Split(string(byteData), "\n")
		for _, line := range data {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("行：%s\t重复次数：%d\n", line, n)
		}
	}
}
