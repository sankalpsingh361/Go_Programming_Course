package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := genn(q)

	receivee(c, q)

	fmt.Println("about to exit")
}

func receivee(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:

			return
		}
	}
}

func genn(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		q<-1
		close(c)
	}()

	return c
}
