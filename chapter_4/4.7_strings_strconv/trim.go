package main

import (
	"strings"
	"fmt"
)

func main() {
	s := " Hi jms "
	space:= strings.TrimSpace(s)
	fmt.Println(s)
	fmt.Println(space)
}
