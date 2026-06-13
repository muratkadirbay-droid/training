package main

import "fmt"

func main() {
	numbers := []int{4, 7, 2, 9, 5}

	numberStatus := make(map[int]string)

	for _, n := range numbers {
		if n%2 == 0 {
			numberStatus[n] = "even"
		} else {
			numberStatus[n] = "odd"
		}
	}

	for num, status := range numberStatus {
		fmt.Println(num, status)
	}
}