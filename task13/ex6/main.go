package main

import "fmt"

func main() {
	firstList := []int{1, 2, 3}
	secondList := []int{1, 2, 4}

	equal := true

	if len(firstList) != len(secondList) {
		equal = false
	} else {
		for i := 0; i < len(firstList); i++ {
			if firstList[i] != secondList[i] {
				equal = false
				break
			}
		}
	}

	if equal {
		fmt.Println("Массивы равны")
	} else {
		fmt.Println("Массивы не равны")
	}
}