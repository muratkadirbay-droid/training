package main

import "fmt"

func squareWorker(n int, ch chan int) {
	ch <- n * n
}

func main() {
	ch := make(chan int)

	go squareWorker(9, ch)

	result := <-ch
	fmt.Println(result)
}