package main

import (
	"fmt"
)

const author string = "Murat"

func main() {
	text := "Go — простой и мощный язык программирования"

	fmt.Printf("Автор %s написал: %s\n", author, text)
	fmt.Println("Длина строки text:", len(text))
}