package main

import (
	"fmt"
	"sync"
)

func checkStatus(site string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Сайт", site, "доступен")
}

func main() {
	var wg sync.WaitGroup

	sites := []string{"google.com", "yandex.ru", "go.dev"}

	for _, s := range sites {
		wg.Add(1)
		go checkStatus(s, &wg)
	}

	wg.Wait()
}