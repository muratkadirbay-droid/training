package main

import (
	"fmt"
	"sync"
	"time"
)

func processData(id int, category string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Запуск процесса %d в категории %s\n", id, category)

	time.Sleep(delay)

	fmt.Printf("Процесс %d завершен\n", id)
}

func main() {
	var wg sync.WaitGroup

	var category string
	fmt.Print("Введите категорию: ")
	fmt.Scan(&category)

	delays := []time.Duration{
		500 * time.Millisecond,
		1 * time.Second,
		200 * time.Millisecond,
	}

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go processData(i, category, delays[i-1], &wg)
	}

	wg.Wait()
}