package main

import "fmt"

func main() {
	var numbers []int
	var x int

	fmt.Println("Ввод чисел (0 = стоп):")

	for {
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		numbers = append(numbers, x)
	}

	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	fmt.Println("Сумма:", sum)
}