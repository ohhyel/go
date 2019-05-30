package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Type()
}

type Rectangle struct {
	width float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64  {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) Type()  {
	fmt.Println("Rectangle")
}

func (r Circle) Type()  {
	fmt.Println("Circle")
}

func main() {
	rec := []Shape{
		Rectangle{width: 10, height: 20},
		Circle{radius: 12},
	}

	for _, s := range rec {
		s.Type()
		fmt.Println("Area: ", s.Area())
	}
}
