package main

import (
	"fmt"
)

func main() {
	var input string

	fmt.Print("Введите строку: ")
	fmt.Scan(&input)

	length := len(input)

	if length < 5 {
		fmt.Println("Слишком короткое слово")
	} else if length <= 10 {
		fmt.Println("Нормальная длина")
	} else {
		fmt.Println("Слишком длинное слово")
	}
}