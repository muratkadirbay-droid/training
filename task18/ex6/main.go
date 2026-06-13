package main

import "fmt"

func checkLogin(login, password string) bool {
	users := map[string]string{
		"admin": "1234",
	}

	if users[login] == password {
		return true
	}
	return false
}

func main() {
	fmt.Println(checkLogin("admin", "1234"))
	fmt.Println(checkLogin("admin", "0000"))
}