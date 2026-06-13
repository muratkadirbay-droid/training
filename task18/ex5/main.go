package main

import "fmt"

func sumSlice(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(sumSlice(arr))
}