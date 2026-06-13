package main

import "fmt"

type Car struct {
	Brand string
	Year  int
}

func main() {
	c := Car{
		Brand: "Toyota",
		Year:  2010,
	}

	ptr := &c
	ptr.Year = 2022

	fmt.Println(c.Brand, c.Year)
}