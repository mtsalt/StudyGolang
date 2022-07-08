package main

import "fmt"

func main() {

	a := map[string]int{
		"x": 100,
		"y": 200,
	}

	fmt.Println(a)
	fmt.Println(a["x"])

	a["z"] = 300
	fmt.Println(a)

	delete(a, "z")

	fmt.Println(a, len(a))

	_, ok := a["z"]
	if ok {
		fmt.Println("exist")
	} else {
		fmt.Println("Not exist")
	}

	for key, value := range a {
		fmt.Printf("%s=%d", key, value)
	}
}
