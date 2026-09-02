package main

import (
	"fmt"
	"time"
)

func CountTransactions() {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Фоновая транзакция №%d обработана\n", i)

		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go CountTransactions()
	fmt.Println("Главный поток main: запуск фоновой работы...")

	time.Sleep(500 * time.Millisecond)
}
