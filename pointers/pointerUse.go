package main

import "fmt"

func main() {
	x := 0
	fmt.Println("x before",&x)
	fmt.Println("x before",x)
	pointerUse(&x)
	fmt.Println("x after",&x)
	fmt.Println("x after",x)
}

func pointerUse(y *int) {
	fmt.Println("y before",y)
	fmt.Println("y before",*y)
	*y = 43
	fmt.Println("y after",y)
	fmt.Println("y after",*y)
}