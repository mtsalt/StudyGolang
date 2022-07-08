package main

import "fmt"

func main() {

	a := [3]int{1, 2, 3}
	fmt.Println(a)

	b := [3]string{}
	b[0] = "a"
	b[1] = "b"
	b[2] = "c"
	fmt.Println(b)

	c := [...]int{3, 4, 5, 6, 7, 8, 9}
	fmt.Println(c)

}
