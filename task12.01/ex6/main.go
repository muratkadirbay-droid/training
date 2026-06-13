package main

import (
	"fmt"
)

func main() {
	var input string

	fmt.Print("Введите строку: ")
	fmt.Scan(&input)

	reversed := ""

	for i := len(input) - 1; i >= 0; i-- {
		reversed += string(input[i])
	}

	fmt.Println("Оригинал:", input)
	fmt.Println("Обратная:", reversed)

	if input == reversed {
		fmt.Println("Это палиндром")
	} else {
		fmt.Println("Это не палиндром")
	}
}