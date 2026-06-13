package main

import "fmt"

func maxNumber(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxNumber(10, 20))
	fmt.Println(maxNumber(50, 15))
}