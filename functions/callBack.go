// passing a function as an argument to another function
package main

import "fmt"

func summ(x, y int) int {
	return x + y
}
func partialSum(x int) func(int) int {
	return func(y int) int {
		return summ(x, y)
	}
}
func main() {
	partial := partialSum(3)
	fmt.Println(partial(7))
}