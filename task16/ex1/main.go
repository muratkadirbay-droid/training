package main

import "fmt"

func main() {
	toolUsage := map[string]int{
		"Go":     3,
		"VSCode": 5,
		"Git":    2,
	}

	for tool, count := range toolUsage {
		fmt.Println(tool, count)
	}
}