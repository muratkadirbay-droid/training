package main

import "fmt"

func main() {
	// Константы
	const BaseRate = 5.50      // цена за 1 кг
	const TaxRate = 0.12       // налог 12%
	const DistanceRate = 2.0   // цена за 1 км
	const FragileFee = 0.2     // наценка 20% за хрупкость

	// Переменные для ввода
	var name string
	var weight float64
	var distance float64
	var fragileCount int

	// Ввод данных
	fmt.Print("Введите имя: ")
	fmt.Scan(&name)

	fmt.Print("Вес груза (кг): ")
	fmt.Scan(&weight)

	fmt.Print("Дистанция (км): ")
	fmt.Scan(&distance)

	fmt.Print("Количество хрупких упаковок: ")
	fmt.Scan(&fragileCount)

	// Расчет базовой стоимости
	baseCost := (weight * BaseRate) * (1 + FragileFee*float64(fragileCount)) + (distance * DistanceRate)

	// Итоговая стоимость с налогом
	totalCost := baseCost * (1 + TaxRate)

	// Вывод результата
	fmt.Println("\n--- Отчет о доставке ---")
	fmt.Println("Отправитель:", name)
	fmt.Printf("Итоговая стоимость: %.2f\n", totalCost)
}