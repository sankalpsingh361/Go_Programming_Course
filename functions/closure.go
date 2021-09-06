package main

import "fmt"

// declaring a variable
var gfg =  0
// assigning a anonymous function to a variable
func main(){
	counter := func() int {
		gfg = gfg + 1
		return gfg
	}
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}
