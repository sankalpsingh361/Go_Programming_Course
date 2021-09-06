package main

import "fmt"

type person2 struct { // new struct which will go to embedded
	first string
	last  string
	age   int
}

type secretAgent struct {
	person2 // previous struct embedded here.!!
	ltk     bool
}

func main() {
	sa1 := secretAgent{
		person2: person2{ // embedded struct
			first: "james",
			last:  "bond",
			age:   34,
		},
		ltk: true,
	}

	p2 := person2{
		first: "rohit",
		last:  "raj",
		age:   24,
	}
	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Println(sa1.ltk, sa1.first, sa1.last, sa1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
