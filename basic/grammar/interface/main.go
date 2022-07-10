package main

import "fmt"

type Printable interface {
	ToString() string
}

func PrintOut(p Printable) {
	fmt.Println(p.ToString())
}

type Person struct {
	name string
}

func (p Person) ToString() string {
	return p.name
}

type Book struct {
	title string
}

func (b Book) ToString() string {
	return b.title
}

func main() {
	a1 := Person{name: "Yamada"}
	a2 := Book{title: "I'm a cat."}
	PrintOut(a1)
	PrintOut(a2)
}
