package main

import "fmt"

func main() {
	var number int
	count := 0

	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	for number > 0 {
		number = number / 10
		count++
	}

	fmt.Println("Количество цифр:", count)
}