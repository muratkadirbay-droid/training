package main

import "fmt"

func main() {
	isLoggedIn := true
	isAdmin := false

	hasAccess := (isLoggedIn && isAdmin) || (isLoggedIn && !isAdmin)

	fmt.Println(hasAccess)
}
