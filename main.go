package main

import (
	"fmt"
)

func SendSingleTx(ch chan string) {
	ch <- "TX_VIP_777"
	close(ch)
}

func main() {
	cryptoChan := make(chan string)
	go SendSingleTx(cryptoChan)
	tx1, ok1 := <-cryptoChan
	fmt.Printf("Чтение 1: Значение: %s, Открыт: %t\n", tx1, ok1)

	tx2, ok2 := <-cryptoChan
	fmt.Printf("Чтение 2: Значение: %s, Открыт: %t\n", tx2, ok2)
}
