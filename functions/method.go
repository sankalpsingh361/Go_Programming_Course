package main

import "fmt"

type person struct {
	first string
	last string
}

type Agent struct {
	person
	ltk bool
}

func (a Agent) mission (){  // method
	fmt.Println("i am ",a.first,a.last)
}

func main() {
	var a = Agent{      // Composite literal
		person: person{
			first: "james",
			last: "bond",
		},
		ltk: true,
	}

	var b = Agent{
		person: person{
			first: "sankalp",
			last: "singh",
		},
		ltk: false,
	}

	fmt.Println(a)
	fmt.Println(b)
	a.mission()
	b.mission()

}