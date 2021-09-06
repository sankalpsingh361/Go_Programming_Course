package main

import "fmt"

// Create a type person struct
type person struct {
	first string
}

// attach a method speak to a type person using a pointer receiver
func (p *person) speak() {
	fmt.Println("hello")
}

// create a type human interface
type human interface {
	speak()
}

// create func "saySomething"
// take human as a parameter
// call the speak method
func saySomething(h human) {
	h.speak()
}

// you can pass a value of type *person into saySomething
func main() {
	p1 := person{
		first: "sankalp",
	}
	saySomething(&p1)

	// does not work
	//saySomething(p1)

	p1.speak()

}
