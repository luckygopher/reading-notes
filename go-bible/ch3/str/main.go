// @Description:字符串练习
// @Author: Arvin
// @date: 2021/3/2 2:47 下午
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(intsToString([]int{1, 2, 3}))
	BytesPackage()
}

// 看似是把系统路径的前缀删除，同时将看似文件类型的后缀名部分删除
func basename1(s string) string {
	for i := len(s); i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

// 将一个表示整数值的字符串，每隔三个字符插入一个逗号分隔符
// 例如“12345”处理后成为“12,345”
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
