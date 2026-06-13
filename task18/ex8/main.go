package main

import "fmt"

func resetAttempts(attempts *int) {
	if *attempts > 3 {
		*attempts = 0
	}
}

func main() {
	attempts := 5

	resetAttempts(&attempts)

	fmt.Println(attempts)
}