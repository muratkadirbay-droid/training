package main

import "fmt"

func main() {
	balance := 3000

	for {
		var choice int

		fmt.Println("1 - Показать баланс")
		fmt.Println("2 - Пополнить +500")
		fmt.Println("3 - Снять -200")
		fmt.Println("0 - Выход")

		fmt.Print("Выберите: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Баланс:", balance)
		}

		if choice == 2 {
			balance += 500
			fmt.Println("Баланс:", balance)
		}

		if choice == 3 {
			balance -= 200
			fmt.Println("Баланс:", balance)
		}

		if choice == 0 {
			fmt.Println("Выход из программы")
			break
		}
	}
}