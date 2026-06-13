package main

import "fmt"

func main() {
	var data []int
	var x int

	fmt.Println("Ввод чисел (0 = стоп):")

	for {
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		data = append(data, x)
	}

	if len(data) > 2 {
		data = append(data[:2], data[3:]...)
	}

	fmt.Println(data)
}