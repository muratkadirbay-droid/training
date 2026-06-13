package main

import "fmt"

func main() {
	myWishList := []string{"Laptop", "Headphones", "Book"}
	friendsWishList := []string{"Watch", "Gamepad"}

	registrationList := []string{}

	for _, item := range myWishList {
		registrationList = append(registrationList, item)
	}

	for _, item := range friendsWishList {
		registrationList = append(registrationList, item)
	}

	fmt.Println(registrationList)
}