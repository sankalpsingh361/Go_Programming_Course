package main

import "fmt"

func main() {
defer 	first()  // defer wala end me run hoga.
	second()
}

func first(){
	fmt.Println("FIRST")
}

func second() {
	fmt.Println("SECOND")
}