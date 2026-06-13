package main

import "fmt"

func main() {
	// 1. Меню (map)
	menu := map[string]float64{
		"Эспрессо": 800,
		"Латте":    1200,
		"Капучино": 1100,
		"Сэндвич":  1500,
		"Круассан": 900,
	}

	// 2. Корзина (slice)
	var order []string

	fmt.Println("===== МЕНЮ =====")
	for name, price := range menu {
		fmt.Printf("%s - %.0f тг\n", name, price)
	}

	// 3. Ввод заказов
	var item string

	for {
		fmt.Print("Введите блюдо (или 'exit'): ")
		fmt.Scan(&item)

		if item == "exit" {
			break
		}

		if _, ok := menu[item]; ok {
			order = append(order, item)
			fmt.Println("Добавлено:", item)
		} else {
			fmt.Println("К сожалению, этого блюда нет в меню")
		}
	}

	// 4. Группировка заказов (map)
	counts := make(map[string]int)

	for _, v := range order {
		counts[v]++
	}

	// 5. Расчёт суммы
	var total float64 = 0

	fmt.Println("\n===== ЧЕК =====")
	fmt.Println("Список покупок:")

	for item, count := range counts {
		price := menu[item] * float64(count)
		total += price
		fmt.Printf("%s x%d — %.0f тг\n", item, count, price)
	}

	// 6. Скидка
	discount := 0.0
	if total > 5000 {
		discount = total * 0.10
	}

	afterDiscount := total - discount

	// 7. НДС 12%
	vat := afterDiscount * 0.12
	final := afterDiscount + vat

	// 8. Итоговый вывод
	fmt.Println("---------------------")
	fmt.Printf("Сумма без скидки: %.0f тг\n", total)
	fmt.Printf("Скидка: %.0f тг\n", discount)
	fmt.Printf("Сумма после скидки: %.0f тг\n", afterDiscount)
	fmt.Printf("НДС (12%%): %.0f тг\n", vat)
	fmt.Printf("ИТОГО К ОПЛАТЕ: %.0f тг\n", final)
}