package main

import "fmt"

func main() {
	examResults := map[string]int{
		"Aruzhan": 85,
		"Dias":    92,
		"Alina":   78,
	}

	for name, score := range examResults {
		if score >= 80 {
			fmt.Println(name, "сдал экзамен")
		} else {
			fmt.Println(name, "не сдал экзамен")
		}
	}
}