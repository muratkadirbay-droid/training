package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "I like Java"

	if strings.Contains(s, "Java") {
		fmt.Println("Найдено")
	} else {
		fmt.Println("Не найдено")
	}
}