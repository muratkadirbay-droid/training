package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)

	if len(s) < 8 {
		fmt.Println("Слишком короткий пароль")
		return
	}

	hasDigit := false
	hasUpper := false

	for _, ch := range s {
		if unicode.IsDigit(ch) {
			hasDigit = true
		}
		if unicode.IsUpper(ch) {
			hasUpper = true
		}
	}

	if !hasDigit {
		fmt.Println("Пароль должен содержать цифру")
	} else if !hasUpper {
		fmt.Println("Пароль должен содержать заглавную букву")
	} else {
		fmt.Println("Пароль корректный")
	}
}