package main

import "fmt"

func sendGreeting(ch chan string) {
	ch <- "Привет из горутины!"
}

func main() {
	ch := make(chan string)

	go sendGreeting(ch)

	msg := <-ch
	fmt.Println(msg)
}