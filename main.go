package main

import (
	"fmt"
	"time"
)

func FetchSlowBybit(ch chan string) {
	time.Sleep(500 * time.Millisecond)
	ch <- "Bybit: Успешный ответ!"
}

func main() {
	bybitChan := make(chan string)
	go FetchSlowBybit(bybitChan)
	fmt.Println("Поток main: Запрашиваю данные у Bybit со встроенной защитой 100 мс...")

	select {
	case res := <-bybitChan:
		fmt.Println("Успех!", res)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("❌ Тайм-аут! Биржа Bybit зависла, сбрасываем соединение.")
	}
}
