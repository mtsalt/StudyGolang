package main

import "fmt"

func main() {

	n := 0
	for n <= 5 {
		fmt.Println(n)
		n++
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}

	nums := [5]int{5, 4, 3, 2, 1}
	for _, i := range nums {
		fmt.Printf("%d ", i)
	}
}
