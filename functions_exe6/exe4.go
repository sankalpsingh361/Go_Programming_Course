package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(" i am ", p.first, p.last, "and ", p.age, "year old")
}
func main() {
	p1 := person{
		first: "sankalp",
		last:  "singh",
		age:   23,
	}

	p1.speak()
}
