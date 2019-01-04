package main

import (
	"fmt"
	"strconv"
	"strings"
)

// StringLean https://books.studygolang.com/gopl-zh/ch3/ch3-05.html
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
	fmt.Printf("length: %d, 16进制: % x\n", len("\U00004e16\U0000754c"), "\U00004e16\U0000754c") // \U 4字节Unicode码点表示法
	fmt.Printf("length: %d, 16进制: % x\n", len("\u4e16\u754c"), "\u4e16\u754c")                 // \u 2字节Unicode码点表示法（常用字符），每个字符都对应一个 Unicode 码点，并不是实际存储的二进制内容。

	fmt.Printf("length: %d, 16进制: % x\n", len("\xe4\xb8\x96\xe7\x95\x8c"), "\xe4\xb8\x96\xe7\x95\x8c") // 用十六进制表示 6个字节, 和二进制存储完全对应。
	fmt.Printf("length: %d, 16进制: % x\n", len("世界"), "世界")                                             // 以十六进制方式输出。展示了以二进制实际存储的内容。

	// rune：存储格式为 int32，语义上相当于一个 Unicode 码点。
	// 将[]rune类型转换应用到UTF8编码的字符串，将返回字符串编码的Unicode码点序列：
	s := "プログラム"
	fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(s)         // 转换为 Unicode 码点序列（整数）
	fmt.Printf("%x\n", r)  // "[30d7 30ed 30b0 30e9 30e0]"

	fmt.Print("-- 数字和字符串转换 --\n")
	x := 123
	y := fmt.Sprintf("%d", x) // 使用格式化字符串
	y1 := strconv.Itoa(x)     // 整数转 ASCII（字符串）
	fmt.Printf("整形转字符串：%T, %T\n", y, y1)
	xx := 133.4
	yy := fmt.Sprintf("%.2f", xx)
	yy1 := strconv.FormatFloat(xx, 'E', -1, 64)
	fmt.Printf("浮点转字符串：%s(%T); %s(%T)\n", yy, yy, yy1, yy1)

	yyy, _ := strconv.ParseInt("1234", 0, 32)
	fmt.Printf("字符串转整形：%d\n", yyy)

	yyyy, _ := strconv.ParseFloat("1234.234", 64)
	fmt.Printf("字符串转浮点：%.3f\n", yyyy)
}
