package main

import "fmt"

func add(x int, y ...int) int {
	ans := x
	for _, i := range y {
		ans += i
	}
	return ans
}

func main() {

	fmt.Printf("ans=%d\n", add(1, 2))
	fmt.Printf("ans=%d\n", add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
