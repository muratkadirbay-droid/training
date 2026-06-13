package main

import "fmt"

func main() {
	var a [11][11]string

	// заполнение точками
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			a[i][j] = "."
		}
	}

	center := 11 / 2

	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			if i == center || j == center || i == j || i+j == 10 {
				a[i][j] = "*"
			}
		}
	}

	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			fmt.Print(a[i][j], " ")
		}
		fmt.Println()
	}
}