package main

import (
	"strings"
	"fmt"
)

func main() {
	s := "Hello world"
	isContains := strings.Contains(s, "world1")
	fmt.Printf("%t", isContains)
}
