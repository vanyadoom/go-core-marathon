package main

import (
	"fmt"
	"time"
)

func FetchBinance(ch chan string) {
	time.Sleep(200 * time.Millisecond)
	ch <- "Binance: BTC = 64500 USD"
}
func FetchBybit(ch chan string) {
	time.Sleep(50 * time.Millisecond)
	ch <- "Bybit: BTC = 64480 USD"
}

func main() {
	binanceChan := make(chan string)
	bybitChan := make(chan string)

	go FetchBinance(binanceChan)
	go FetchBybit(bybitChan)

	fmt.Println("Поток main: Запускаю опрос бирж наперегонки...")

	select {
	case res := <-binanceChan:
		fmt.Println("Победил агрегатор!", res)
	case res := <-bybitChan:
		fmt.Println("Победил агрегатор!", res)
	}
}
