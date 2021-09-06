package main

import "fmt"

func main() {
	c := make(chan int)

	// put 10 number in the channel
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	// pull the numbers and print them
	for v := range c {
		fmt.Println(v)

	}
	fmt.Println("about to exit")
}
