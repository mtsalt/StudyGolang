package main

import (
	"fmt"
	"time"
)

func funcA() {
	for i := 0; i < 10; i++ {
		fmt.Print("A")
		time.Sleep(10 * time.Millisecond)
	}
}

func funcB(chB chan<- string) {
	time.Sleep(3 * time.Second)
	chB <- "Finished"
}

func main() {
	go funcA()
	for i := 0; i < 10; i++ {
		fmt.Print("B")
		time.Sleep(20 * time.Millisecond)
	}
	fmt.Println()

	chB := make(chan string)
	defer close(chB)
	go funcB(chB)
	msg := <-chB
	fmt.Println(msg)
}
