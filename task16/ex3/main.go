package main

import "fmt"

func main() {
	var name string
	fmt.Scan(&name)

	userInfo := map[string]string{
		"name":       name,
		"isLoggedIn": "true",
	}

	fmt.Println("Пользователь", userInfo["name"], "вошёл в систему")
}