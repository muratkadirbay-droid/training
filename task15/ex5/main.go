package main

import "fmt"

func main() {
	var words []string
	var s string

	fmt.Println("Ввод слов (stop = конец):")

	for {
		fmt.Scan(&s)
		if s == "stop" {
			break
		}
		words = append(words, s)
	}

	var reversed []string

	for i := len(words) - 1; i >= 0; i-- {
		reversed = append(reversed, words[i])
	}

	fmt.Println("Оригинал:", words)
	fmt.Println("Перевёрнутый:", reversed)
}