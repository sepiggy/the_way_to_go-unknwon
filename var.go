package main

import (
	"fmt"
	"os"
)

// 因式分解法声明变量
var (
	a   int
	b   bool
	str string
)

// 在编译时, 自动推断类型
var (
	d    = 22
	e    = true
	city = "Beijing"
)

// 在运行时, 自动推断类型 (包级别的全局变量)
var (
	HOME   = os.Getenv("HOME")
	USER   = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)

func main() {
	a = 15
	b = true
	str = "Michael"
	fmt.Println(a, b, str)
	fmt.Println(d, e, city)
	fmt.Println(HOME, USER, GOROOT)
}
