package main

import "fmt"

func main() {
	number := 5

	for i := 1; i <= 10; i++ {
		fmt.Println(number, "*", i, "=", number*i)
	}
}