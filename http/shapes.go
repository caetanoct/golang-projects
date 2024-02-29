package main

import (
	"fmt"
)

type triangle struct{
	height float64
	base float64
}
type square struct{
	sideLength float64
}

type shape interface{
	area() float64
}

func main() {
	t := triangle{4.0, 2}
	s := square{3.0}
	printArea(s)
	printArea(t)
}

func printArea(s shape) {
	fmt.Println(s.area())
}
func (t triangle) area() float64 {
	return (t.base * t.height) / 2 
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}
