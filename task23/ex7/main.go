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

	delays := make([]time.Duration, 3)

	for i := 0; i < 3; i++ {
		var seconds int
		fmt.Printf("Введите задержку для процесса %d (в секундах): ", i+1)
		fmt.Scan(&seconds)
		delays[i] = time.Duration(seconds) * time.Second
	}

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go processData(i, category, delays[i-1], &wg)
	}

	wg.Wait()
}