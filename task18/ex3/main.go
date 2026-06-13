package main

import "fmt"

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	fmt.Println(isEven(2))
	fmt.Println(isEven(7))
	fmt.Println(isEven(10))
}