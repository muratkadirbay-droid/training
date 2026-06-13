package main

import "fmt"

func main() {
	sales := [][]int{
		{10, 20, 30},
		{5, 15, 25},
	}

	total := make(map[int]int)

	for i := 0; i < len(sales); i++ {
		sum := 0
		for j := 0; j < len(sales[i]); j++ {
			sum += sales[i][j]
		}
		total[i] = sum
	}

	for shop, sum := range total {
		fmt.Println("Магазин", shop, ":", sum)
	}
}