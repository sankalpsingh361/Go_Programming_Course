package main

import "fmt"

type person struct {
	name string
}

func main() {
	p1 := person{
		name:"sankalp",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)

}
func changeMe (p *person)  {
	p.name = "shanky singh"
	//(*p).name = "shanty's"
}