package main

import (
	"fmt"
	"strings"
)

func main() {
	println("main")
	// s := []int{23, 44, 5, 7}
	// println(s[2])

	fmt.Printf("%q\n", strings.Split("a,b,c", ""))
	fmt.Printf("%d\n", 1993%10)
	fmt.Printf("%d\n", 1993/10)

	fmt.Print("----\n")
	fmt.Printf("%T\n", "ssfrfff"[0])
	fmt.Printf("%T\n", "ssfrfff"[0:3])
}
