package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "Mama went to the market and bought apples"

	countA := 0

	for i := 0; i < len(sentence); i++ {
		ch := sentence[i]
		if ch == 'a' || ch == 'A' {
			countA++
		}
	}

	fmt.Printf("В строке %d символов и %d букв а\n", len(sentence), countA)

	_ = strings.Contains(sentence, "a") // просто пример использования strings
}