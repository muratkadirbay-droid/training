package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func calc(c Circle) {
	area := math.Pi * c.Radius * c.Radius
	circ := 2 * math.Pi * c.Radius

	fmt.Println("Area:", area)
	fmt.Println("Circumference:", circ)
}

func main() {
	c := Circle{Radius: 5}

	calc(c)

	c.Radius = 10
	calc(c)
}