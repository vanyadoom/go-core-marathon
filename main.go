package main

import (
	"fmt"
	"time"
)

func GeneralSignal(ch chan string) {

	time.Sleep(300 * time.Millisecond)

	ch <- "BTC"

}

func main() {

	signalChan := make(chan string)

	go GeneralSignal(signalChan)

	fmt.Println("Главный поток main: Робот ИИ запущен, жду данные из канала...", signalChan)

	token := <-signalChan

	fmt.Printf("Главный поток main: Сигнал получен! Срочно покупаем актив: %s\n", token)

}
