package main

import (
	"fmt"
)

func main() {
	var text string

	fmt.Print("Введите строку: ")
	fmt.Scan(&text)

	length := len(text)

	words := 1
	vowels := 0

	vowelSet := "aeiouyAEIOUY"

	for i := 0; i < len(text); i++ {
		ch := text[i]

		// подсчет слов
		if ch == ' ' {
			words++
		}

		// подсчет гласных
		for j := 0; j < len(vowelSet); j++ {
			if ch == vowelSet[j] {
				vowels++
			}
		}
	}

	fmt.Printf("Длина строки: %d, слов: %d, гласных: %d\n", length, words, vowels)
}