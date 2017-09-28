package main

import (
	"fmt"
	"runtime"
)

var str string = ""

func main() {

	if str == "" {
		fmt.Println("This string is null")
	}

	if len(str) == 0 {
		fmt.Println("This string is null too")
	}

	if runtime.GOOS == "darwin" {
		fmt.Println("This is a Mac OS System")
	}

	fmt.Println(Abs(-100))
	fmt.Println(isGreater(100, 99))

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isGreater(x, y int) bool {
	if x > y {
		return true
	}
	return false
}
