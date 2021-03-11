// @Description: 字符串处理：bytes、strings、strconv、unicode包使用
// @Author: Arvin
// @date: 2021/3/3 2:07 下午
package main

import (
	"bytes"
	"fmt"
)

// bytes包
func BytesPackage() {
	var (
		s  = []byte("hello world")
		ss = []byte("hello")
	)
	// 将 s 中的所有字符修改为大写(小写、标题)格式返回
	_ = bytes.ToUpper(s)
	_ = bytes.ToLower(s)
	_ = bytes.ToTitle(s)

	// 将 s 中的所有单词的首字符修改为 Title 格式返回
	// BUG：不能很好的处理以 Unicode 标点符号分割的单词
	_ = bytes.Title(s)

	// 比较两个 []byte，nil 参数相当于空 []byte。
	// s <  ss 返回 -1
	// s == ss 返回 0
	// s >  ss 返回 1
	_ = bytes.Compare(s, ss)

	// 判断 s、ss 是否相等，nil 参数相当于空 []byte
	_ = bytes.Equal(s, ss)

	// 判断 s、ss 是否相似，忽略大写、小写、标题三种格式的区别
	_ = bytes.EqualFold(s, ss)

	// 去掉 s 两边（左边、右边）包含在 cutset 中的字符(返回 s 的切片)
	_ = bytes.Trim(s, "cutset")
	_ = bytes.TrimLeft(s, "cutset")
	_ = bytes.TrimRight(s, "cutset")
	// 去掉 s 两边的空白（unicode.IsSpace）(返回 s 的切片)
	_ = bytes.TrimSpace(s)
	// 去掉 s 的前缀 prefix（后缀suffix）（返回 s 的切片）
	_ = bytes.TrimPrefix(s, ss)
	_ = bytes.TrimSuffix(s, ss)

	// 如果 ss 为空，则将 s 切分成 Unicode 字符列表
	// 以 ss 为分隔符将 s 切分成多个子串，结果不包含分隔符
	_ = bytes.Split(s, ss)
	// 指定切分次数 n，超出 n 的部分将不进行切分
	_ = bytes.SplitN(s, ss, 2)
	// 功能同 Split，只不过结果包含分隔符(在各个子串尾部)
	_ = bytes.SplitAfter(s, ss)
	_ = bytes.SplitAfterN(s, ss, 2)
	// 以连续空白为分隔符将 s 切分成多个子串，结果不包含分隔符
	_ = bytes.Fields(s)
	// 以 ss 为连接符，将子串列表 s 连接成一个字节串
	_ = bytes.Join([][]byte{s, s}, ss)
	// 将子串 s 重复 count 次后返回
	bytes.Repeat(s, 10)

	// 判断 s 是否有前缀 ss（后缀 ss）
	bytes.HasPrefix(s, ss)
	bytes.HasSuffix(s, ss)
	// 判断 s 中是否包含子串 ss（字符 r）
	bytes.Contains(s, ss)
	bytes.ContainsRune(s, 'r')
	// 判断 s 中是否包含 chars 中的任何一个字符
	bytes.ContainsAny(s, "chars")

	// 查找子串 ss（字节 c、字符 r）在 s 中第一次出现的位置，找不到则返回 -1
	bytes.Index(s, ss)
	bytes.IndexByte(s, 'c')
	bytes.IndexRune(s, 'r')
	// 查找 chars 中的任何一个字符在 s 中第一次出现的位置，找不到则返回 -1
	bytes.IndexAny(s, "chars")
	// 功能同上，只不过查找最后一次出现的位置
	bytes.LastIndex(s, ss)
	bytes.LastIndexByte(s, 'c')
	bytes.LastIndexAny(s, "chars")
	// 获取 ss 在 s 中出现的次数（ss 不能重叠）
	bytes.Count(s, ss)

	// 将 s 中前 n 个 old 替换为 new，n < 0 则替换全部
	old := []byte{97}
	new := []byte{100}
	bytes.Replace(s, old, new, 2)
	// 将 s 转换为 []rune 类型返回
	bytes.Runes(s)

	// Reader
	b1 := []byte("Hello World!")
	b2 := []byte("Hello 世界！")
	buf := make([]byte, 6)
	// 将 b1 包装成 bytes.Reader 对象
	rd := bytes.NewReader(b1)
	// 读取数据到buf
	rd.Read(buf)
	fmt.Printf("%q\n", buf) // "Hello "
	rd.Read(buf)
	fmt.Printf("%q\n", buf) // "World!"
	// 将底层数据切换为 b2，同时复位所有标记
	rd.Reset(b2)
	rd.Read(buf)
	// "Hello "
	fmt.Printf("%q\n", buf)
	// 底层数据的总长度 Size:15 未读取部分的数据长度 Len:9
	fmt.Printf("Size:%d, Len:%d\n", rd.Size(), rd.Len())

	// Buffer
	obj := bytes.NewBufferString("Hello World!")
	buff := make([]byte, 6)
	// 获取数据切片
	bb := obj.Bytes()
	// 读出一部分数据，看看切片有没有变化
	obj.Read(buff)
	fmt.Printf("%s\n", obj.String()) // World!
	fmt.Printf("%s\n\n", bb)         // Hello World!

	// 写入一部分数据，看看切片有没有变化
	obj.Write([]byte("abcdefg"))
	fmt.Printf("%s\n", obj.String()) // World!abcdefg
	fmt.Printf("%s\n\n", bb)         // Hello World!

	// 再读出一部分数据，看看切片有没有变化
	obj.Read(buff)
	fmt.Printf("%s\n", obj.String()) // abcdefg
	fmt.Printf("%s\n", bb)           // Hello World!
}
