package main

import (
	"strings"
	"fmt"
)

func main() {
	str := "Hello Jms Hello Jms Hello Jms"
	replace := strings.Replace(str, "Jms", "Tom", -1)
	fmt.Println(replace)
}
