package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	aCount := 0
	eCount := 0

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if ch == 'a' || ch == 'A' {
			aCount++
		}
		if ch == 'e' || ch == 'E' {
			eCount++
		}
	}

	fmt.Println("A/a:", aCount)
	fmt.Println("E/e:", eCount)
}