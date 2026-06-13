package main

import (
	"fmt"
)

const (
	BaseTariff    = 0.45 // цена за 1 кВт·ч
	HighLoadTax   = 0.15 // 15%
	NightDiscount = 0.30 // 30%
)

func main() {
	for {
		var name string

		fmt.Print("Введите название прибора (или 'done'): ")
		fmt.Scan(&name)

		if name == "done" {
			fmt.Println("Расчет завершен.")
			break
		}

		var power float64
		var hours float64
		var night bool

		fmt.Print("Мощность (Вт): ")
		fmt.Scan(&power)

		fmt.Print("Время работы (ч): ")
		fmt.Scan(&hours)

		fmt.Print("Ночной режим (true/false): ")
		fmt.Scan(&night)

		// Расчет энергии
		consumption := (power * hours) / 1000.0
		cost := consumption * BaseTariff

		// Категория мощности
		category := ""
		switch {
		case power < 100:
			category = "Экономный"
		case power >= 100 && power <= 1000:
			category = "Стандартный"
		default:
			category = "Мощный"
		}

		// Ночной режим (скидка)
		if night {
			cost = cost * (1 - NightDiscount)
		}

		// Высокая нагрузка (налог)
		if consumption > 10 {
			cost = cost * (1 + HighLoadTax)
		}

		// Вывод отчёта
		fmt.Println("\n--- Отчет по прибору ---")
		fmt.Printf("Прибор: %s [Категория: %s]\n", name, category)
		fmt.Printf("Расход: %.2f кВт·ч\n", consumption)
		fmt.Printf("Итоговая стоимость: %.2f (с учетом налогов/скидок)\n", cost)
		fmt.Println("------------------------\n")
	}
}
