package main

import "fmt"

func main() {
	var age int

	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)

	if age < 13 {
		fmt.Println("Ребёнок")
	} else if age <= 17 {
		fmt.Println("Подросток")
	} else {
		fmt.Println("Взрослый")
	}
}