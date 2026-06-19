package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Регистрируем обработчик
	http.HandleFunc("/calculate", calculateHandler)

	fmt.Println("Server is listening on port 8080...")
	
	// Запускаем сервер
	http.ListenAndServe(":8080", nil)
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	// Переменные
	var a, b float64
	var op string

	// Читаем данные из тела запроса
	_, err := fmt.Fscan(r.Body, &a, &b, &op)

	if err != nil {
		fmt.Fprintf(w, "Ошибка! Отправьте данные в формате: 10 20 +")
		return
	}

	var result float64

	// Логика калькулятора
	switch op {
	case "+":
		result = a + b

	case "-":
		result = a - b

	case "*":
		result = a * b

	case "/":
		// Проверка деления на ноль
		if b == 0 {
			fmt.Fprintf(w, "Ошибка: Деление на ноль")
			return
		}
		result = a / b

	default:
		fmt.Fprintf(w, "Неизвестная операция: %s", op)
		return
	}

	// Отправляем результат
	fmt.Fprintf(w, "Результат: %f", result)
}