package main

import "fmt"

type (
	IZ  int
	FZ  float64
	STR string
)

func main() {
	// 类型转换
	a := 5.0
	var b int
	b = int(a)
	fmt.Println(b)
}
