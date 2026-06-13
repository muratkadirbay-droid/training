package main

import (
	"errors"
	"fmt"
	"math"
)

// ------------------- Задание 1 -------------------
func task1() {
	products := map[string]int{
		"Apple":  100,
		"Banana": 70,
		"Orange": 120,
	}

	name := "Banana"

	price, ok := products[name]
	if ok {
		fmt.Println("Цена товара:", price)
	} else {
		fmt.Println("Товар отсутствует")
	}
}

// ------------------- Задание 2 -------------------
func divide(a int, b int) (int, interface{}) {
	if b == 0 {
		return 0, "Ошибка: деление на ноль невозможно"
	}
	return a / b, nil
}

func task2() {
	result, err := divide(10, 2)
	fmt.Println("10 / 2 =", result, "error:", err)

	result, err = divide(10, 0)
	fmt.Println("10 / 0 =", result, "error:", err)
}

// ------------------- Задание 3 -------------------
func checkAge(age int) error {
	if age < 0 || age > 120 {
		return errors.New("указан некорректный возраст")
	}
	return nil
}

func task3() {
	ages := []int{25, -5, 130}

	for _, age := range ages {
		err := checkAge(age)
		if err != nil {
			fmt.Println("Возраст", age, "ошибка:", err)
		} else {
			fmt.Println("Возраст", age, "корректный")
		}
	}
}

// ------------------- Задание 4 -------------------
func validatePassword(password string) (bool, error) {
	if len(password) < 6 {
		return false, errors.New("пароль слишком короткий")
	}
	return true, nil
}

func task4() {
	pass1 := "12345"
	pass2 := "secure123"

	for _, p := range []string{pass1, pass2} {
		ok, err := validatePassword(p)
		fmt.Println("Пароль:", p, "OK:", ok, "Error:", err)
	}
}

// ------------------- Задание 5 -------------------
type invalidRadiusError struct{}

func (e invalidRadiusError) Error() string {
	return "радиус не может быть отрицательным"
}

func task5() {
	var err error = invalidRadiusError{}
	fmt.Println(err.Error())
}

// ------------------- Задание 6 -------------------
func calculateCircleArea(radius int) (float64, error) {
	if radius < 0 {
		return 0, invalidRadiusError{}
	}
	return math.Pi * float64(radius) * float64(radius), nil
}

func task6() {
	area, err := calculateCircleArea(-5)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Площадь:", area)
	}
}

// ------------------- Задание 7 -------------------
func findUser(users []string, name string) (int, error) {
	for i, u := range users {
		if u == name {
			return i, nil
		}
	}
	return -1, errors.New("пользователь не найден")
}

func task7() {
	users := []string{"Alice", "Bob", "Charlie"}

	index, err := findUser(users, "Bob")
	fmt.Println("Bob:", index, err)

	index, err = findUser(users, "David")
	fmt.Println("David:", index, err)
}

// ------------------- main -------------------
func main() {
	fmt.Println("----- Task 1 -----")
	task1()

	fmt.Println("\n----- Task 2 -----")
	task2()

	fmt.Println("\n----- Task 3 -----")
	task3()

	fmt.Println("\n----- Task 4 -----")
	task4()

	fmt.Println("\n----- Task 5 -----")
	task5()

	fmt.Println("\n----- Task 6 -----")
	task6()

	fmt.Println("\n----- Task 7 -----")
	task7()
}