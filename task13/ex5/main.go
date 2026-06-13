package main

import "fmt"

func main() {
	friendsList := []string{"Aruzhan", "Dias", "Bekbolat", "Alikhan"}

	found := false

	for _, name := range friendsList {
		if name == "Bekbolat" {
			found = true
			break
		}
	}

	if found {
		fmt.Println("Мне очень повезло")
	} else {
		fmt.Println("Мне не повезло")
	}
}