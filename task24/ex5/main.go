package main

import "fmt"

func filterEven(input chan int, output chan int) {
	for v := range input {
		if v%2 == 0 {
			output <- v
		}
	}
	close(output)
}

func main() {
	input := make(chan int)
	output := make(chan int)

	go filterEven(input, output)

	for i := 1; i <= 10; i++ {
		input <- i
	}
	close(input)

	for v := range output {
		fmt.Println(v)
	}
}