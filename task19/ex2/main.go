package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func main() {
	b := Book{
		Title:  "Go Basics",
		Author: "Robert",
		Pages:  250,
	}

	fmt.Println(b.Title)
	fmt.Println(b.Author)
	fmt.Println(b.Pages)
}