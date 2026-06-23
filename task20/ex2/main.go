package main

import "fmt"

type Student struct {
	Name   string
	Grades []int
}

func (s Student) Average() float64 {
	sum := 0
	for _, g := range s.Grades {
		sum += g
	}
	return float64(sum) / float64(len(s.Grades))
}

func (s Student) GetName() string {
	return s.Name
}

type Person interface {
	Average() float64
	GetName() string
}

func main() {
	s1 := Student{"Aruzhan", []int{80, 90, 85}}
	s2 := Student{"Dias", []int{70, 75, 80}}
	s3 := Student{"Alina", []int{95, 90, 100}}

	people := []Person{s1, s2, s3}

	for _, p := range people {
		fmt.Printf("Студент %s: средний балл %.2f\n", p.GetName(), p.Average())
	}
}