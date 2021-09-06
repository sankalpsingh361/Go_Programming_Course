package main

import "fmt"

func main() {
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "sankalp",
		friends: map[string]int{
			"sankalp": 23,
			"shanky":  34343,
			"rohit": 5685,
		},
		favDrinks: []string{
			"coca-cola",
			" Mango shake",
		},
	}
	fmt.Println(s)
	fmt.Println("------------------")

	for k, v := range s.friends{
		fmt.Println(k,v)
	}
	for i, v := range s.favDrinks{
		fmt.Println(i,v)
	}
}
