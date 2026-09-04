package main

import (
	"fmt"
)

func main() {

	queue := make(chan string, 2)

	queue <- "TX_MARKET"
	queue <- "TX_LIMIT"

	fmt.Printf("Очередь заполнена! Вместимость: %d, Текущая длина: %d\n", cap(queue), len(queue))

	tx1 := <-queue

	fmt.Println("Обработана транзакция:", tx1)

	tx2 := <-queue

	fmt.Println("Обработана транзакция:", tx2)

}
