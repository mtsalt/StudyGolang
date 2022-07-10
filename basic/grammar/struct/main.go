package main

import (
	"fmt"
	"struct/person"
)

func main() {

	pl := person.NewPersonList()

	fmt.Println("### Add item to PersonList. ###")

	p1 := person.NewPerson(1, "John", 25)
	pl.Add(p1)

	p2 := person.NewPerson(2, "Bob", 24)
	pl.Add(p2)

	p3 := person.NewPerson(3, "Mike", 28)
	pl.Add(p3)

	pl.PrintList()

	fmt.Println("### Remove item from PersonList. ###")
	pl.Remove(1) // remove John
	pl.Remove(2) // remove Bob
	pl.Remove(3) // remove Mike
	pl.PrintList()
}
