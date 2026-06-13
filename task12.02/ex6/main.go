package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Go is fast and Go is simple"

	count := strings.Count(s, "Go")
	fmt.Println(count)
}