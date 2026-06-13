package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	found := false

	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] && s[i+1] == s[i+2] {
			found = true
			break
		}
	}

	if found {
		fmt.Println("Есть тройка одинаковых символов")
	} else {
		fmt.Println("Нет тройки")
	}
}