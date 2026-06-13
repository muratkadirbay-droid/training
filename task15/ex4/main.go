package main

import "fmt"

func main() {
	var temps []int
	var x int

	fmt.Println("Ввод температур (0 = стоп):")

	for {
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		temps = append(temps, x)
	}

	min := temps[0]
	max := temps[0]

	for i := 0; i < len(temps); i++ {
		if temps[i] < min {
			min = temps[i]
		}
		if temps[i] > max {
			max = temps[i]
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}