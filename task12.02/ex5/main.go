package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	if strings.HasPrefix(s, "admin") {
		fmt.Println("Доступ разрешён")
	} else {
		fmt.Println("Доступ запрещён")
	}
}