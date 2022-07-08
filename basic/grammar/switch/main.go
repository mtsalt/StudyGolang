package main

import "fmt"

func main() {

	num := 1
	switch num {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
	default:
		fmt.Println(2)
	}
}
