package main

import "fmt"

func main() {
	words := []string{"go", "is", "fast"}

	wordLength := make(map[string]int)

	for _, w := range words {
		wordLength[w] = len(w)
	}

	for word, length := range wordLength {
		fmt.Println(word, length)
	}
}