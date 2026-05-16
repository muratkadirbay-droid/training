package main

import "fmt"

func main() {

	// Задание 1
	schooling := 11
	fmt.Println("Years of schooling:", schooling)

	schooling = schooling + 1
	fmt.Println("After extra year:", schooling)

	// Задание 2
	name := "Vladislav"
	fmt.Println("Name:", name)

	name = "Murat"
	fmt.Println("My real name:", name)

	// Задание 3
	steps := 0
	fmt.Println("Steps in the morning:", steps)

	steps = 2000
	fmt.Println("Steps at noon:", steps)
	fmt.Println("Хорошая работа! Вы уже на пути к своей цели")

	//  Задание 4
	largeNumber := 6000000
	fmt.Println("Large number:", largeNumber)

	//  Задание 5
	const breakTime = 15
	fmt.Println("Break time:", breakTime)

	// breakTime = 20
	// Нельзя изменить константу, потому что const — это неизменяемый константа.
}
