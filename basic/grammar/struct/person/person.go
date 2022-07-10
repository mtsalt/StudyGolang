package person

type Person struct {
	id   int
	name string
	age  int
}

func NewPerson(id int, name string, age int) Person {
	return Person{id: id, name: name, age: age}
}

func (p *Person) Person() (id int, name string, age int) {
	return p.id, p.name, p.age
}
