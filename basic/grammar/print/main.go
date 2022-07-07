package main

import (
	"fmt"
)

func main() {

	num := 1
	str := "abc"

	fmt.Print("num=", num, " str=", str)
	fmt.Println("num=", num, "str=", str)
	fmt.Printf("num=%d str=%s", num, str)

	// Documentation > https://pkg.go.dev/fmt
}
