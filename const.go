package main

import "fmt"

const Pi = 3.14159

// 常量并行赋值
const beef, two, tomato = "eat", 2, "veg"
const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6
const (
	Jan = 1
	Feb = 2
	Mar = 3
)

// 枚举
const (
	UnKnown = 0
	Female  = 1
	Male    = 2
)

const (
	a1 = iota
	a2
	a3
)

var n int

func main() {
	fmt.Println(Pi)

	n = 10
	fmt.Println(n + 5)

	fmt.Println(UnKnown, Female, Male)

	fmt.Println(a1, a2, a3)

}
