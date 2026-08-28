package main

import (
	"fmt"
)

func main() {

	fmt.Println("Действие 1: Подключение к кошельку")

	defer fmt.Println("Действие 3: Безопасное отключение от кошелька")

	fmt.Println("Действие 2: Проведение транзакции...")

}
