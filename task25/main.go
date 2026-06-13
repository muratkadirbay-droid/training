package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

/* =======================
   1. КОНСТАНТЫ
======================= */

const (
	BaseDeliveryRate   = 1.0
	MaxBatteryLevel    = 100.0
	EmergencyThreshold = 20.0
	BufferCapacity     = 5
)

/* =======================
   2. ТИПЫ РОБОТОВ
======================= */

type RobotType string

const (
	Drone        RobotType = "Drone"
	Wheeled      RobotType = "Wheeled"
	HeavyLifter  RobotType = "HeavyLifter"
)

/* =======================
   3. СТРУКТУРЫ
======================= */

type Product struct {
	ID     int
	Name   string
	Weight float64
}

type Robot struct {
	ID          int
	Model       string
	Type        RobotType
	Battery     float64
	IsAvailable bool
}

/* =======================
   4. ИНТЕРФЕЙС
======================= */

type Mover interface {
	Move(distance float64) error
}

/* =======================
   5. MOVE (метод Robot)
======================= */

func (r *Robot) Move(distance float64) error {

	var coeff float64

	switch r.Type {
	case Drone:
		coeff = 1.5
	case Wheeled:
		coeff = 0.8
	case HeavyLifter:
		coeff = 2.5
	default:
		coeff = 1.0
	}

	required := distance * coeff

	if r.Battery < required {
		return errors.New("недостаточно энергии")
	}

	r.Battery -= required
	return nil
}

/* =======================
   6. RECHARGE
======================= */

func Recharge(r *Robot) {
	r.Battery = MaxBatteryLevel
}

/* =======================
   7. DELIVERY TASK (GOROUTINE)
======================= */

func DeliveryTask(r *Robot, p Product, distance float64, ch chan<- string) {

	time.Sleep(time.Duration(distance*100) * time.Millisecond)

	err := r.Move(distance)

	status := "успешно"

	if err != nil {
		status = "ошибка"
	}

	report := fmt.Sprintf(
		"Робот %d доставил %s | статус: %s",
		r.ID, p.Name, status,
	)

	ch <- strings.TrimSpace(strings.ToUpper(report))
}

/* =======================
   8. MAIN
======================= */

func main() {

	/* склад */
	warehouse := map[string]int{
		"Apple":  3,
		"Book":   2,
		"Phone":  1,
		"Mouse":  2,
	}

	/* роботы */
	robots := []Robot{
		{1, "RX-1", Drone, 100, true},
		{2, "WX-2", Wheeled, 80, true},
		{3, "HL-9", HeavyLifter, 60, true},
	}

	/* буфер заказов */
	var orderBuffer [BufferCapacity]Product
	orderCount := 0
	nextID := 1

	for {

		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Добавить товар")
		fmt.Println("2. Склад")
		fmt.Println("3. Доставка")
		fmt.Println("4. Выход")

		var choice int
		fmt.Scan(&choice)

		/* =======================
		   1. ДОБАВИТЬ В ЗАКАЗ
		======================= */
		if choice == 1 {

			if orderCount >= BufferCapacity {
				fmt.Println("Буфер заказов полон!")
				continue
			}

			var name string
			fmt.Print("Товар: ")
			fmt.Scan(&name)

			if warehouse[name] <= 0 {
				fmt.Println("Товара нет на складе")
				continue
			}

			warehouse[name]--

			orderBuffer[orderCount] = Product{
				ID:     nextID,
				Name:   name,
				Weight: 1.0,
			}

			nextID++
			orderCount++

			fmt.Println("Добавлено в заказ")

		}

		/* =======================
		   2. СКЛАД
		======================= */
		if choice == 2 {

			fmt.Println("\n--- Склад ---")
			for k, v := range warehouse {
				fmt.Println(k, ":", v)
			}

			fmt.Println("\n--- Роботы ---")
			for _, r := range robots {
				status := "свободен"
				if !r.IsAvailable {
					status = "занят"
				}
				fmt.Printf("Robot %d (%s): Battery %.1f, %s\n",
					r.ID, r.Type, r.Battery, status)
			}

			fmt.Println("\n--- Буфер ---")
			for i := 0; i < orderCount; i++ {
				fmt.Printf("%d: %s\n",
					orderBuffer[i].ID,
					orderBuffer[i].Name)
			}
		}

		/* =======================
		   3. ДОСТАВКА
		======================= */
		if choice == 3 {

			if orderCount == 0 {
				fmt.Println("Нет заказов")
				continue
			}

			reportChan := make(chan string)
			activeJobs := 0

			for i := 0; i < orderCount; i++ {

				product := orderBuffer[i]

				var chosen *Robot

				for j := range robots {

					if robots[j].IsAvailable &&
						robots[j].Battery > EmergencyThreshold {

						// проверка веса (упрощённо)
						if robots[j].Type == Drone && product.Weight > 2 {
							continue
						}

						chosen = &robots[j]
						break
					}
				}

				if chosen == nil {
					fmt.Println("Нет робота для:", product.Name)
					continue
				}

				chosen.IsAvailable = false
				activeJobs++

				go DeliveryTask(chosen, product, 2.0, reportChan)
			}

			for i := 0; i < activeJobs; i++ {
				msg := <-reportChan
				fmt.Println(msg)
			}

			// восстановление
			for i := range robots {
				if robots[i].Battery < EmergencyThreshold {
					Recharge(&robots[i])
				}
				robots[i].IsAvailable = true
			}

			orderCount = 0
		}

		/* =======================
		   4. EXIT
		======================= */
		if choice == 4 {
			fmt.Println("Завершение работы...")
			break
		}
	}
}