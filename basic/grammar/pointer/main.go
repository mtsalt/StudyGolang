package main

import "fmt"

func ponterF(value *int) {
	*value = 100
}

func F(value int) {
	value = 100
}

func main() {

	var a int
	var p *int

	p = &a
	*p = 1

	fmt.Println(a)

	var b, c int
	ponterF(&b)
	F(c)

	fmt.Println(b)
	fmt.Println(c)
}
