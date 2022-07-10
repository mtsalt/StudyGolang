package person

import "fmt"

type PersonList struct {
	personList []Person
}

func NewPersonList() PersonList {
	return PersonList{}
}

func (pl *PersonList) Add(p Person) {
	pl.personList = append(pl.personList, p)
}

func (pl *PersonList) Remove(id int) bool {

	targetIndex := -1
	for i, p := range pl.personList {
		if p.id == id {
			targetIndex = i
		}
	}

	if targetIndex != -1 {
		pl.personList = append(pl.personList[:targetIndex], pl.personList[targetIndex+1:]...)
		return true
	} else {
		return false
	}
}

func (pl *PersonList) PrintList() {
	for _, p := range pl.personList {
		fmt.Printf("%#v\n", p)
	}
}
