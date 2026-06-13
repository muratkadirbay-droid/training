package main

import (
	"fmt"
	"time"
)

func printInfo() {
	fmt.Println("Горутина запущена")
}

func main() {
	go printInfo()

	time.Sleep(time.Second)
}