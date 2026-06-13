package main

import "fmt"

func main() {
	var a [4][4]int

	// ввод массива
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	var i, j int
	fmt.Scan(&i, &j)

	// swap строк
	for k := 0; k < 4; k++ {
		a[i][k], a[j][k] = a[j][k], a[i][k]
	}

	// вывод
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(a[i][j], " ")
		}
		fmt.Println()
	}
}