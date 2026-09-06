package main

import (
	"fmt"
	"sync"
)

var balance int
var mu sync.Mutex
var wg sync.WaitGroup

func Deposit() {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	balance++
}

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go Deposit()
	}

	wg.Wait()
	fmt.Println("Юбилейный баланс биржи равен:", balance)

}
