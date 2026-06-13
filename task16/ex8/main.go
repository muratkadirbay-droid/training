package main

import "fmt"

func main() {
	loginAttempts := map[string]int{
		"admin": 2,
		"guest": 0,
	}

	loginAttempts["admin"]++

	if loginAttempts["admin"] > 2 {
		fmt.Println("Доступ заблокирован")
	}
}