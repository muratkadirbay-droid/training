package main

import (
	"fmt"
)

// ------------------- Задание 1 -------------------
func printData() {
	fmt.Println("Обработка данных")
}

func task1() {
	fmt.Println("Старт программы")
	defer printData()
	fmt.Println("Конец программы")
}

// ------------------- Задание 2 -------------------
func task2() {
	defer fmt.Println("Первый")
	defer fmt.Println("Второй")
	defer fmt.Println("Третий")
}

// ------------------- Задание 3 -------------------
func checkAge(age int) {
	if age < 0 {
		panic("Возраст не может быть отрицательным!")
	}
	fmt.Println("Возраст:", age)
}

func task3() {
	checkAge(-5)
}

// ------------------- Задание 4 -------------------
func cleanup() {
	fmt.Println("Очистка ресурсов выполнена")
}

func task4() {
	defer cleanup()
	panic("Критическая ошибка!")
}

// ------------------- Задание 5 -------------------
func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Паника успешно перехвачена:", r)
	}
}

func task5() {
	defer handlePanic()

	panic("Что-то пошло не так!")
}

// ------------------- Задание 6 -------------------
func safeDivide(a, b float64) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Паника перехвачена:", r)
		}
	}()

	if b == 0 {
		panic("Деление на ноль!")
	}

	fmt.Println("Результат:", a/b)
}

func task6() {
	safeDivide(10, 0)
	fmt.Println("Программа продолжает работу после safeDivide")
}

// ------------------- Задание 7 -------------------
func getElement(slice []int, index int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Паника перехвачена (out of bounds):", r)
		}
	}()

	fmt.Println("Элемент:", slice[index])
}

func task7() {
	nums := []int{10, 20, 30}
	getElement(nums, 10)
	fmt.Println("Программа продолжает работу после getElement")
}

// ------------------- main -------------------
func main() {
	fmt.Println("----- Task 1 -----")
	task1()

	fmt.Println("\n----- Task 2 -----")
	task2()

	fmt.Println("\n----- Task 3 -----")
	// task3() // раскомментируй, чтобы увидеть panic

	fmt.Println("\n----- Task 4 -----")
	// task4() // раскомментируй, чтобы увидеть defer + panic

	fmt.Println("\n----- Task 5 -----")
	task5()

	fmt.Println("\n----- Task 6 -----")
	task6()

	fmt.Println("\n----- Task 7 -----")
	task7()
}