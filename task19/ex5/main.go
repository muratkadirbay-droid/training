package main

import "fmt"

type Student struct {
	Name  string
	Grade int
}

func main() {
	s1 := Student{"Ali", 85}
	s2 := Student{"Dias", 92}

	fmt.Printf("Студент %s получил %d баллов\n", s1.Name, s1.Grade)
	fmt.Printf("Студент %s получил %d баллов\n", s2.Name, s2.Grade)

	if s1.Grade > s2.Grade {
		fmt.Println("Лучший:", s1.Name)
	} else {
		fmt.Println("Лучший:", s2.Name)
	}
}