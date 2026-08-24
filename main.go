package main

import "fmt"

func main() {

	balances := make(map[string]float64)

	balances["USD"] = 500.0

	amount, ok := balances["EUR"]

	if ok {

		fmt.Printf("Кошелёк найден! Баланс: %.2f\n", amount)

	} else {

		fmt.Printf("Ошибка: Запрошенный валютный кошелек не существует!")
	}

}
