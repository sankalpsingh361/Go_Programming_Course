package main

import "fmt"

func main(){
	a := cb()
	fmt.Println(a())
}

func cb() func() int{   // callback function
	return func() int {
		return 23
	}
}