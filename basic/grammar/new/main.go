package main

import "fmt"

type Book struct {
	Title string
	Autor string
	Isbn  string
}

func main() {

	b := new(Book)
	b.Title = "The Art of Readable Code"
	b.Autor = "Dustin Boswell and Trevor Foucher"
	b.Isbn = "978-4-87311-565-8"

	fmt.Printf("%#v\n", b)
}
