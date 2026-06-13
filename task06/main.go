package main

import "fmt"

func main() {

	age := 28
	fmt.Print(age)
	age = age + 1
	fmt.Print(age)

	height := 170
	height_in_meters := 1.7

	fmt.Print(height)
	fmt.Print(height_in_meters)

	isStudent := true
	fmt.Println(isStudent)

	temperature := 25
	fmt.Println(temperature)

	if temperature >= 20 {
		fmt.Println("Погода теплая")
	} else {
		fmt.Println("Погода холодная")
	}

	favoriteQuote := "Не ошибается тот, кто ничего не делает."
	fmt.Println(favoriteQuote)

	const PI = 3.14
	fmt.Println(PI)
}