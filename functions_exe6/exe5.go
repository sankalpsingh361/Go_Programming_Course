package main
import (
	"fmt"
	"math"
)
type square struct {
	length float64
}
type circle struct {
	radius float64
}
func (s square) area() float64 {
	return s.length * s.length
}
func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
type shape interface {
	area() float64
}
func info(s shape) {
	x := s.area()
	fmt.Println(x)
}
func main() {
	circ := circle{5.5}
	squa := square{4.1}

	info(circ)
	info(squa)
}
