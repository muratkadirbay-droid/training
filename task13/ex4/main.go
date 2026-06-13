package main

import "fmt"

func main() {
	numbersList := []int{1, 2, 3, 4, 5}

	found := false

	for i := 0; i < len(numbersList); i++ {
		if numbersList[i] == 3 {
			found = true
			break
		}
	}

	if found {
		fmt.Println("Число 3 есть в массиве")
	} else {
		fmt.Println("Число 3 отсутствует в массиве")
	}
}