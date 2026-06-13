package main

import "fmt"

func sumReader(ch chan int) {
	sum := 0
	for v := range ch {
		sum += v
	}
	fmt.Println("Сумма:", sum)
}

func main() {
	ch := make(chan int)

	go sumReader(ch)

	ch <- 10
	ch <- 20
	ch <- 30

	close(ch)
}