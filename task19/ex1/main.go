package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{
		Name: "Aruzhan",
		Age:  20,
	}

	fmt.Println(p.Name)
	fmt.Println(p.Age)
}