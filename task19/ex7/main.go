package main

import "fmt"

type Car struct {
	Brand string
	Speed int
}

func CheckSpeed(car Car) string {
	if car.Speed > 100 {
		return "Слишком быстро"
	} else if car.Speed >= 60 && car.Speed <= 100 {
		return "Нормальная скорость"
	} else {
		return "Медленно"
	}
}

func main() {
	car1 := Car{"Toyota", 120}
	car2 := Car{"BMW", 80}
	car3 := Car{"Lada", 40}

	fmt.Printf("Машина %s: %s\n", car1.Brand, CheckSpeed(car1))
	fmt.Printf("Машина %s: %s\n", car2.Brand, CheckSpeed(car2))
	fmt.Printf("Машина %s: %s\n", car3.Brand, CheckSpeed(car3))
}