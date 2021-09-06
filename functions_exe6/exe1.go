package main

import "fmt"

func foo() int {
	return 23

}

func bar() (int, string) {

 return 1998, " My DOB "
}

func main() {
	a := foo()
	b,c := bar()

	fmt.Println(a,b,c)
}