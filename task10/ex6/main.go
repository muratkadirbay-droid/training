package main

import "fmt"

func main() {
	balance := -10

	if balance >= 0 {
		fmt.Println("Баланс положительный")
	} else {
		fmt.Println("Баланс отрицательный")
	}
}