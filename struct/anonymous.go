package main

import "fmt"

func main() {
	p1 := struct {   // P1 = Composite Literal
		first string
		last string
		age int
	}{
		first: "sankalp",
		last: "singh",
		age: 23,
	}
	fmt.Println(p1)  // Print Struct.
	fmt.Println(p1.first, p1.last, p1.age)   // Values of Struct.

}
