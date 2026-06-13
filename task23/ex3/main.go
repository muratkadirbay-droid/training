package main

import (
	"fmt"
	"sync"
)

func heavyTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Выполняю задачу...")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go heavyTask(&wg)

	wg.Wait()
}