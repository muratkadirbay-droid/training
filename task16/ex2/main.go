package main

import "fmt"

func main() {
	buildStatus := map[string]bool{
		"build": true,
		"run":   false,
	}

	if buildStatus["build"] {
		fmt.Println("Сборка прошла успешно")
	}
}