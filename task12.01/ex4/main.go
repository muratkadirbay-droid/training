package main

import (
	"fmt"
)

func main() {
	word := "Hello"

	fmt.Println("Первый символ:", string(word[0]))
	fmt.Println("Последний символ:", string(word[len(word)-1]))
	fmt.Println("Длина строки:", len(word))

	for i := 0; i < len(word); i++ {
		fmt.Println(string(word[i]))
	}
}