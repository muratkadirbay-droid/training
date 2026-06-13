package main

import "fmt"

func main() {
	menu := map[string]int{
		"Burger": 2500,
		"Pizza":  3200,
		"Tea":    500,
	}

	var dish string
	fmt.Scan(&dish)

	price, ok := menu[dish]

	if ok {
		fmt.Println(price)
	} else {
		fmt.Println("Блюдо не найдено")
	}
}