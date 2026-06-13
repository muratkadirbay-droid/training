package main

import "fmt"

func main() {
	toyList := []string{"Car", "Doll", "Ball"}

	testToyList := make([]string, len(toyList))
	copy(testToyList, toyList)

	testToyList[1] = "Boat"

	fmt.Println("Оригинал:", toyList)
	fmt.Println("Копия:", testToyList)
}