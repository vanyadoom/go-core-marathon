package main

import (
	"fmt"
	"time"
)

func StreamTransation(ch chan string) {
	for i := 1; i <= 3; i++ {
		ch <- fmt.Sprintf("TX_ID_%d", i)
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func main() {
	txChan := make(chan string)
	go StreamTransation(txChan)
	for tx := range txChan {
		fmt.Println("Главный поток принял из канала:", tx)
	}
	fmt.Println("Поток main: Канал успешно закрылся, конвейер завершён!")
}
