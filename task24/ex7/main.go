package main

import "fmt"

func main() {
	ch := make(chan string, 3)

	ch <- "Первое сообщение"
	ch <- "Второе сообщение"
	ch <- "Третье сообщение"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// ❗ Попытка отправить 4-е сообщение без чтения
	// ch <- "Четвёртое" // приведёт к deadlock (зависание)
}