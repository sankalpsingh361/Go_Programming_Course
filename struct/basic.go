package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func main() {
	p1 := person{
		first: "sankalp",
		last: "singh",
		age: 23,
	}
	p2 := person{
		first: "shanky",
		last: "singh",
		age: 22,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last , p1.age)
	fmt.Println(p2.first, p2.last , p2.age)
}
