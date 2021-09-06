// closure also provide another aspect which is data isolation
package main

import "fmt"

// New Counter Function to Isolate Global variable
func newCounter() func() int {
	var gfg = 0
	return func() int {
		gfg = gfg + 1
		return gfg
	}

}
func main() {
	// newCounter Function is assigned to a variable
	counter := newCounter()
	// invoke counter
	fmt.Println(counter())
	fmt.Println(counter())
}
