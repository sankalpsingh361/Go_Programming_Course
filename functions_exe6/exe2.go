// Variadic parameter

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	x := fooo(nums...)
	fmt.Println(x)

}

func fooo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
