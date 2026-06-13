package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func main() {
	r1 := Rectangle{Width: 5, Height: 3}
	r2 := Rectangle{Width: 10, Height: 2}

	shapes := []Shape{r1, r2}

	for i, s := range shapes {
		fmt.Printf("Прямоугольник %d: площадь %.2f, периметр %.2f\n",
			i+1, s.Area(), s.Perimeter())
	}
}