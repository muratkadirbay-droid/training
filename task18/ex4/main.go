package main

import "fmt"

func greetUser(name string) {
	fmt.Println("Привет,", name)
}

func main() {
	greetUser("Aruzhan")
	greetUser("Dias")
}