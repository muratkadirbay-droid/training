package main

import "fmt"

func main() {
	var myWishList []string
	var friendsWishList []string

	var s string

	fmt.Println("Мои желания (stop):")
	for {
		fmt.Scan(&s)
		if s == "stop" {
			break
		}
		myWishList = append(myWishList, s)
	}

	fmt.Println("Желания друга (stop):")
	for {
		fmt.Scan(&s)
		if s == "stop" {
			break
		}
		friendsWishList = append(friendsWishList, s)
	}

	var registrationList []string

	for _, v := range myWishList {
		registrationList = append(registrationList, v)
	}
	for _, v := range friendsWishList {
		registrationList = append(registrationList, v)
	}

	fmt.Println("Общий список:", registrationList)
}