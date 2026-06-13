package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	s = strings.ToLower(s)

	reversed := ""

	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}

	if s == reversed {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Не палиндром")
	}
}