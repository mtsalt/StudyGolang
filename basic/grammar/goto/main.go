package main

import "fmt"

func main() {

	n := 100
	switch n {
	case 1:
		goto Label1
	case 10:
		goto Label10
	case 100:
		goto Label100
	}
Label1:
	fmt.Println(1)
	return
Label10:
	fmt.Println(10)
	return
Label100:
	fmt.Println(100)
	return
}
