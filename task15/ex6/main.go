package main

import "fmt"

func main() {
	var nums []int
	var x int

	fmt.Println("Ввод чисел (0 = стоп):")

	for {
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		nums = append(nums, x)
	}

	sorted := true

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			sorted = false
			break
		}
	}

	fmt.Println(sorted)
}