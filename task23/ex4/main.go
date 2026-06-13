package main

import (
	"fmt"
	"sync"
)

func count(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		fmt.Printf("Горутина %d: %d\n", id, i)
	}
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go count(i, &wg)
	}

	wg.Wait()
}