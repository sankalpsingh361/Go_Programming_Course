package main

import "fmt"

func main() {
	defer fou()
	barr()
}

func fou() {
	fmt.Println("i am fou")
}
func barr() {
	fmt.Println("i am barr")
}
