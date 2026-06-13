package main

import "fmt"

func square(n int) int {
	return n * n
}

func main() {
	fmt.Println(square(2))
	fmt.Println(square(5))
	fmt.Println(square(10))
}