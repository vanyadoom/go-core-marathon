package main

import (
	"fmt"
	"time"
)

func CryptoWorker(id int, jobs <-chan int, results chan<- int) {
	for idTx := range jobs {
		fmt.Printf("Воркер №%d: Начал обработку транзакции TX_ID_%d\n", id, idTx)
		time.Sleep(50 * time.Millisecond)
		results <- idTx * 10
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for i := 1; i <= 3; i++ {
		go CryptoWorker(i, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for m := 1; m <= 5; m++ {
		res := <-results
		fmt.Println("Главный поток принял результат комиссии:", res)
	}
}
