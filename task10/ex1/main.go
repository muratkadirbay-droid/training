package main

import "fmt"

func main() {
	temperature := 15 // от -20 до 40

	if temperature < 0 {
		fmt.Println("Холодно")
	} else if temperature <= 20 {
		fmt.Println("Тепло")
	} else {
		fmt.Println("Жарко")
	}
}
