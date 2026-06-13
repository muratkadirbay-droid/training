package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func calc(r Rectangle) {
	area := r.Width * r.Height
	perim := 2 * (r.Width + r.Height)

	fmt.Println("Area:", area)
	fmt.Println("Perimeter:", perim)
}

func main() {
	r := Rectangle{Width: 5, Height: 3}

	calc(r)

	r.Width = 10
	r.Height = 6

	calc(r)
}