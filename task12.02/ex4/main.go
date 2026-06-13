package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "one,two,three"

	result := strings.ReplaceAll(s, ",", ";")
	fmt.Println(result)
}