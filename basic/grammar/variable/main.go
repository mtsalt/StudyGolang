package main

import (
	"fmt"
)

func main() {

	var num int
	var str string

	var num2 int = 123
	var str2 string = "abc"

	num3 := 456
	str3 := "def"

	num4, str4 := 789, "xyz"

	var (
		a int
		b int
		c int
	)

	fmt.Printf("num=%d, str=%s\n", num, str)
	fmt.Printf("num=%d, str=%s\n", num2, str2)
	fmt.Printf("num=%d, str=%s\n", num3, str3)
	fmt.Printf("num=%d, str=%s\n", num4, str4)
	fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)
}
