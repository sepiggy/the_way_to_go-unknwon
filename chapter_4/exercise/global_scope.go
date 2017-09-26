package main

var a1 = "G"

func main() {
	n1()
	m1()
	n1()
}

func n1() {
	print(a1)
}

func m1() {
	a1 = "0"
	print(a1)
}
