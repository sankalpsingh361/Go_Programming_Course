package main

import "fmt"

func newCounter() func() int {
	gfg := 0
	return func() int {
		gfg += 1
		return gfg
	}
}
func main() {
	counter := newCounter()
	fmt.Println(counter())
	fmt.Println(counter())
}

