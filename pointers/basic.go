package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)        // in which memory "address" the keyword "a" saved.
	fmt.Printf("%T\n", a)  // type of value
	fmt.Printf("%T\n", &a) // type of address
	fmt.Println("-----------")

	b := &a  // "a" ke "address" ki value "b" me copy.
	//var b *int = &a
	//var b int = &a // error { int ko nhi, pointer int ko done }

	fmt.Println(b)    // New Address of "b"
	fmt.Println(*b)   // value stored at an address when you have the address.
	fmt.Println(*&a)  // "a" ke "address" pe jo value save hai. = "dereferences"


	*b = 50   // update the value which is stored in existing address.
	fmt.Println(a)
}
