// @Description: 文件操作
// @Author: Arvin
// @date: 2021/3/29 11:06 上午
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// FileOne()
	// FileThree()
	// FileFour()
	FileFive()
}

// 打开文件
func FileOne() {
	// file 可以叫（file对象 or file指针 or file文件句柄）
	file, err := os.Open("./text.txt")
	if err != nil {
		fmt.Println("open file err ", err)
	}
	fmt.Printf("%v", file)
	// 关闭文件
	if err = file.Close(); err != nil {
		fmt.Println("close file err ", err)
	}
}

// 读取文件的内容并显示在终端（带缓冲区的方式）
// 方式1：一行一行的读
func FileTwo() {
	// 打开文件，得到文件句柄
	file, err := os.Open("./text.txt")
	if err != nil {
		fmt.Printf("open file err %v", err)
	}
	// 函数结束之前关闭文件句柄，否则会导致内存泄露
	defer file.Close()
	// 创建带缓冲的 *Reader
	reader := bufio.NewReader(file)
	for {
		// 每读到一个换行符就结束
		str, err := reader.ReadString('\n')
		fmt.Print(str)
		// 是否为文件末尾
		if err == io.EOF {
			break
		}
	}
	fmt.Println("文件已读完")
}

// 方式2：一次性将整个文件读入到内存中
func FileThree() {
	content, err := ioutil.ReadFile("./text.txt")
	if err != nil {
		fmt.Printf("read file err %v", err)
	}
	// 把读取的内容显示到终端
	// []byte切片，会输出一串数字
	fmt.Printf("%v \n", content)
	// 转为字符串输出
	fmt.Printf("%s", string(content))
}

// 写入文件方式1： 创建一个新文件，写入内容
func FileFour() {
	const NAME = "go\n"
	file, err := os.OpenFile("./go.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err %v", err)
	}
	// 关闭文件句柄
	defer file.Close()
	// 创建一个缓冲写入
	writer := bufio.NewWriter(file)
	for i := 0; i < 3; i++ {
		_, err := writer.WriteString(NAME)
		if err != nil {
			fmt.Printf("write err %v", err)
			return
		}
	}
	// 上面只是写入了缓冲区，还需要flush才能写入到真正的文件中
	if err := writer.Flush(); err != nil {
		fmt.Printf("flush err %v", err)
	}
	return
}

// 写入文件方式2
func FileFive() {
	content := "写入内容\n"
	if err := ioutil.WriteFile("./go.txt", []byte(content), 0666); err != nil {
		fmt.Printf("write file err %v", err)
	}
}
