package main

import "fmt"

func increaseBalance(balance *int, amount int) {
	*balance += amount
}

func main() {
	balance := 1000

	increaseBalance(&balance, 500)
	fmt.Println(balance)

	increaseBalance(&balance, 200)
	fmt.Println(balance)
}