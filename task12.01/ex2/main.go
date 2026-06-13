package main

import (
	"fmt"
)

func main() {
	message := "Он сказал: \"Привет!\""

	if message == "" {
		fmt.Println("Строка пуста")
	} else {
		fmt.Println("Строка не пустая")
	}

	fmt.Println("Длина строки:", len(message))
}