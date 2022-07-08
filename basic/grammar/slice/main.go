package main

import "fmt"

func main() {

	a := []int{}
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	fmt.Println(a)

	b := []int{}
	for i := 0; i < 10; i++ {
		b = append(b, i)
		fmt.Println(len(b), cap(b))
	}

	buf := make([]byte, 0, 1024)
	fmt.Println(buf)
}
