package main

import "fmt"

func Convert() {
	// https://www.cnblogs.com/zrtqsk/p/4157350.html

	fmt.Print("-- 类型转换 --\n")
	// 语法：<结果类型> := <目标类型> ( <表达式> )
	// 类型转换是用来在不同但相互兼容的类型之间的相互转换的方式，所以，当类型不兼容的时候，是无法转换的。如下：
	var i int = 99
	v1 := float32(i)
	var v2 int64 = int64(i)
	var v3 = string(i)
	// var4 := []int8(i)  // 不兼容的类型
	fmt.Printf("%T -> %v\n", i, i)
	fmt.Printf("%T -> %v\n", v1, v1)
	fmt.Printf("%T -> %v\n", v2, v2)
	fmt.Printf("%T -> %v\n", v3, v3)

	//创建一个int变量，并获得它的指针
	var i1 = new(int32)
	fmt.Printf("%T -> %v -> %v\n", i1, i1, *i1)
	*i1 = 88
	fmt.Printf("%T -> %v -> %v\n", i1, i1, *i1)
	var v4 = (*int32)(i1)
	fmt.Printf("%T -> %v -> %v\n", v4, v4, *v4)

	fmt.Print("-- 类型断言 --\n")
	// 语法
	// <目标类型的值>，<布尔参数> := <表达式>.( 目标类型 ) // 安全类型断言
	// <目标类型的值> := <表达式>.( 目标类型 )　　//非安全类型断言

	// 类型断言的本质，跟类型转换类似，都是类型之间进行转换，不同之处在于，类型断言实在接口之间进行，相当于Java中，对于一个对象，把一种接口的引用转换成另一种。
	var i2 interface{} = "99"
	// j := i2.(int)  // 不安全的类型断言，类型不匹配会报 panic
	// fmt.Printf("%T->%d\n", j, j)
	j, b := i2.(int) // 安全的类型断言。
	if b {
		fmt.Printf("%T->%d\n", j, j)
	} else {
		fmt.Println("类型不匹配")
	}
}
