package main

import "fmt"

func main()  {
	x,y := greeting("sankalp","singh")
	fmt.Println(x)
	fmt.Println(y)
}

func greeting (fn string, ln string) (string, bool){
 a := fmt.Sprint(fn," " ,ln ," ",`says "hello"`)
 b := false
 return a,b
}
