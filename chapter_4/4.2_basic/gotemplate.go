package main

import "fmt"

const c = "C"

var v int = 5

type T struct {
}

func init() {

}

func main() {
	var a int
	Func1()
	// ...
	fmt.Println(a)
}

func (t T) Method1()  {
	// ...
}

func Func1() {
	// ...
}

