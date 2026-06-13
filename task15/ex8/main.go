package main

import "fmt"

func main() {
	toyList := [3]string{"Car", "Doll", "Ball"}

	var testToyList [3]string
	testToyList = toyList // копирование массива

	testToyList[1] = "Boat"

	fmt.Println("Оригинал:", toyList)
	fmt.Println("Копия:", testToyList)
}