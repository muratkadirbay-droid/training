package main

import (
	"fmt"
	"sync"
)

func calculateSquare(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(num * num)
}

func main() {
	numbers := []int{2, 4, 6, 8}

	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go calculateSquare(num, &wg)
	}

	wg.Wait()
}