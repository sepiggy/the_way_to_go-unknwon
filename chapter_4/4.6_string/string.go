package main

import "fmt"

func main() {
	str := "Beginning of the string " + "second part of the string"
	s := "hel" + "lo,"
	s += " world!"
	fmt.Println(str)
	fmt.Println(s)
}
