package main

import (
	"fmt"
	"time"
)

func sayHello(name string) {
	fmt.Println("Привет,", name)
}

func main() {
	go sayHello("Ali")
	go sayHello("Dias")
	go sayHello("Aruzhan")

	time.Sleep(time.Second)
}