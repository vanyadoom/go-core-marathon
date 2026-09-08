package main

import (
	"fmt"
	"sync"
)

func GenerateQuotes(name string, ch chan string) {
	for i := 1; i <= 2; i++ {
		ch <- fmt.Sprintf("%s: BTC = 64000", name)
	}
	close(ch)
}

func FanIn(channels ...chan string) <-chan string {
	out := make(chan string, 10)
	var wg sync.WaitGroup

	output := func(c <-chan string) {
		for val := range c {
			out <- val
		}
		wg.Done()
	}
	for _, c := range channels {
		wg.Add(1)
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go GenerateQuotes("Binance", ch1)
	go GenerateQuotes("Bybit", ch2)

	merged := FanIn(ch1, ch2)
	for msg := range merged {
		fmt.Println("Агрегатор поймал:", msg)
	}
}
