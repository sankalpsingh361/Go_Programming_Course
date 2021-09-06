package main

import "fmt"

func main() {
	foo()
	bar("james") // Argument
	s1:= woo("sankalp")
	fmt.Println(s1)
}

func foo() {
	fmt.Println("hello from foo")
}

// Everything in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("hello", s)
}

func woo (st string) string {
	return fmt.Sprint("hello from woo ", st)
}