package main

import "fmt"

func main() {
	text := "Developer"

	for _, char := range text {
		fmt.Println(string(char))
	}
}