package main

import (
	"fmt"
)

func main() {

	const a = 100
	const (
		b = 20
		c = "abc"
	)
	fmt.Printf("a=%d, b=%d, c=%s\n", a, b, c)
}
