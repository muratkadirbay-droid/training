package main

import "fmt"

func main() {
	age := 20
	hasTicket := true

	canEnter := age >= 18 && hasTicket

	fmt.Println(canEnter)
}