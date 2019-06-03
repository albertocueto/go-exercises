package main

import (
	"fmt"
	"math"
)

type square struct {
	large float64
}
type circle struct {
	radius float64
}
type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius
}

func (s square) area() float64 {
	return s.large * s.large
}

func info(s shape) {
	switch s.(type) {
	case circle:
		fmt.Println("Circle area: ", s.(circle), s.(circle).area())
	case square:
		fmt.Println("Square area:", s.(square), s.(square).area())
	}
}

func main() {
	circle1 := circle{
		radius: 2,
	}
	square1 := square{
		large: 4,
	}
	info(circle1)
	info(square1)
}
