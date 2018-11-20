package main

import (
	"fmt"

	"github.com/weineel/go-sprint/kit"
)

func main() {
	fmt.Println("main")
	s := []int{23, 44, 5, 7}
	fmt.Println(kit.ReverseInt(s))
}
