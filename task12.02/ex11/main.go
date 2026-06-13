package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)

	first := rune(s[0])
	last := rune(s[len(s)-1])

	ok := true

	if !unicode.IsUpper(first) {
		fmt.Println("Не с заглавной буквы")
		ok = false
	}

	if last != '.' {
		fmt.Println("Не заканчивается точкой")
		ok = false
	}

	if ok {
		fmt.Println("Строка оформлена правильно")
	}
}