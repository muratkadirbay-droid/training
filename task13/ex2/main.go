package main

import "fmt"

func main() {
	subjectsList := []string{"Физика", "Химия", "География"}

	fmt.Println("Первый:", subjectsList[0])
	fmt.Println("Последний:", subjectsList[len(subjectsList)-1])

	subjectsList[1] = "Биология"

	fmt.Println(subjectsList)
}