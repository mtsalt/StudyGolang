package main

type PersonList struct {
	personList []Person
}

type Person struct {
	id   int
	name string
	age  int
}

func (p *Person) SetPerson(id int, name string, age int) {
	p.id = id
	p.name = name
	p.age = age
}

func (p *Person) Person() (int, string, int) {
	return p.id, p.name, p.age
}

func (pl *PersonList) Add(p Person) {
	pl.personList = append(pl.personList, p)
}

func (pl *PersonList) Remove(id int) bool {

	for _, p := range pl.personList {
		if p.id == id {

		}
	}
	return true
}

func main() {
}
