package main

import "fmt"

func main() {
	var command string

	fmt.Print("Введите команду: ")
	fmt.Scan(&command)

	switch command {
	case "start":
		fmt.Println("Запуск")
	case "stop":
		fmt.Println("Остановка")
	case "restart":
		fmt.Println("Перезапуск")
	default:
		fmt.Println("Неизвестная команда")
	}
}