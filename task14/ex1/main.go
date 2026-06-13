package main

import "fmt"

func main() {
	var arr [3][5]int

	// ввод
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scan(&arr[i][j])
		}
	}

	maxVal := arr[0][0]
	maxI, maxJ := 0, 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			if arr[i][j] > maxVal {
				maxVal = arr[i][j]
				maxI = i
				maxJ = j
			}
		}
	}

	fmt.Println(maxI, maxJ)
}