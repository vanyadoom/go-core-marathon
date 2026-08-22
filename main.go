package main

import "fmt"

func ApplyDiscount(amount, discount float64) (float64, bool) {
	if discount < 0 || discount > 100.0 {
		return amount, false
	}
	finalPrice := amount - (amount * (discount / 100.0))
	return finalPrice, true
}

func main() {
	price1, ok1 := ApplyDiscount(5000.0, 15.0)
	price2, ok2 := ApplyDiscount(2500.0, 120.0)

	if ok1 {
		fmt.Printf("Покупка 1. Купон успешно применен! Финальная цена: %.2f\n", price1)
	} else {
		fmt.Printf("Покупка 1. Ошибка: Неверный процент скидки! Цена без изменений: %.2f\n", price1)
	}
	if ok2 {
		fmt.Printf("Покупка 2. Купон успешно применен! Финальная цена: %.2f\n", price2)
	} else {
		fmt.Printf("Покупка 2. Ошибка: Неверный процент скидки! Цена без изменений: %.2f\n", price2)
	}
}
