package main

import (
	"fmt"
	"strings"
)

func StringLean() {
	fmt.Print("-- 字符串操作包 strings --\n")
	fmt.Printf("%q\n", strings.Split("a,b,c", ""))
	fmt.Printf("%d\n", 1993%10)
	fmt.Printf("%d\n", 1993/10)

	fmt.Print("-- 切片 --\n")
	fmt.Printf("%T\n", "ssfrfff"[0])
	fmt.Printf("%T\n", "ssfrfff"[0:3])

	// Unicode 编码，https://books.studygolang.com/gopl-zh/ch3/ch3-05.html
	fmt.Print("-- 以下都是相同的字面值 --\n")
	/*
		Unicode 是一种编码规范，或者说是一种给符号编号的表，每个字符都在这个表中有唯一的编号，Unicode码点就是这个编号。
		utf 是 Unicode 码点和二进制存储之间的转换关系（格式）。

		例子："\U00004e16\U0000754c"、"\u4e16\u754c"、"\xe4\xb8\x96\xe7\x95\x8c"、"世界" 在 utf-8 下是等价的。
		"\U00004e16\U0000754c" 和 "\u4e16\u754c" 都是Unicode码点表示，不管是 utf-8 还是 还是 utf-16 ，都是一种转换格式，都能解析出相同的结果。
		而 "\xe4\xb8\x96\xe7\x95\x8c" 是二进制格式，和二进制存储完全对应，不同的转换格式，表示的字符不同。
		"世界" 是可读字符表示，不同的转换格式，存储的二进制不同。
	*/
	// go 默认是 utf-8
	fmt.Printf("%d\n", len("\U00004e16\U0000754c")) // \U 4字节Unicode码点表示法
	fmt.Printf("%d\n", len("\u4e16\u754c"))         // \u 2字节Unicode码点表示法（常用字符），每个字符都对应一个 Unicode 码点，并不是实际存储的二进制内容。

	fmt.Printf("%d\n", len("\xe4\xb8\x96\xe7\x95\x8c")) // 用十六进制表示 6个字节, 和二进制存储完全对应。
	fmt.Printf("%d\n", len("世界"))
	fmt.Printf("16进制：%x\n", "世界") // 以十六进制方式输出。展示了以二进制实际存储的内容。
}
