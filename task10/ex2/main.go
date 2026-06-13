package main

import "fmt"

func main() {
	score := 85

	if score >= 90 {
		fmt.Println("Отлично")
	} else if score >= 70 {
		fmt.Println("Хорошо")
	} else if score >= 50 {
		fmt.Println("Удовлетворительно")
	} else {
		fmt.Println("Не сдал")
	}
}