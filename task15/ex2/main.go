package main

import "fmt"

func main() {
	var values []int
	var x int

	fmt.Println("Ввод чисел (0 = стоп):")

	for {
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		values = append(values, x)
	}

	var even []int

	for i := 0; i < len(values); i++ {
		if values[i]%2 == 0 {
			even = append(even, values[i])
		}
	}

	fmt.Println("Исходный:", values)
	fmt.Println("Чётные:", even)
}