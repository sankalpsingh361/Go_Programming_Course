package main

import "fmt"

type person1 struct {
	firstName  string
	lastName   string
	favFlavour []string
}

func main() {
	p1 := person1{
		firstName:  "sankalp",
		lastName:   "singh",
		favFlavour: []string{"vanilla", "mango", "chocolate"},
	}
	p2 := person1{
		firstName:  "shanky",
		lastName:   "singh",
		favFlavour: []string{"abc", "def", "ghi"},
	}

	fmt.Println(p1)
	fmt.Println(p1.firstName, p1.lastName, p1.favFlavour)

	fmt.Println(p2)
	fmt.Println(p2.firstName, p2.lastName, p2.favFlavour)

	for i, v := range p1.favFlavour {
		fmt.Println(i, v)
	}
	for i, v := range p2.favFlavour {
		fmt.Println(i, v)
	}
}
