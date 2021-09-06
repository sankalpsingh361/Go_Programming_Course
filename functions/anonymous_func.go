package main

import "fmt"

//func main() {
//	func() {   // anonymous function
//		fmt.Println("hello anonymous function")
//	}()
//}



// Assigning an anonymous func to a variable
//func main(){
//	value := func(){
//		fmt.Println("hello anonymous function")
//	}
//	value()
//}


// Passing argument
func main(){
	func(argu string){
		fmt.Println("passing argument to anonymous function")
	}("text")
}