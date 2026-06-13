package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	atCount := 0
	atIndex := -1

	for i := 0; i < len(s); i++ {
		if s[i] == '@' {
			atCount++
			atIndex = i
		}
	}

	if atCount == 0 {
		fmt.Println("Email должен содержать @")
		return
	}

	if atCount > 1 {
		fmt.Println("Email должен содержать @")
		return
	}

	hasDot := false
	for i := atIndex; i < len(s); i++ {
		if s[i] == '.' {
			hasDot = true
			break
		}
	}

	if !hasDot {
		fmt.Println("Email должен содержать точку после @")
	} else {
		fmt.Println("Email корректен")
	}
}