package main

import "fmt"

func checkChannel(ch chan string) {
	val, ok := <-ch

	if ok {
		fmt.Println("Получено:", val)
	} else {
		fmt.Println("Канал закрыт!")
	}
}

func main() {
	// открытый канал
	ch1 := make(chan string, 1)
	ch1 <- "Hello"
	checkChannel(ch1)

	// закрытый канал
	ch2 := make(chan string, 1)
	close(ch2)
	checkChannel(ch2)
}