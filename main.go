package main

import (
	"fmt"
	"sync"
	"time"
)

func ProcessPayment(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Горутина: Начинаю валидацию платежа...")

	time.Sleep(200 * time.Millisecond)

	fmt.Println("Горутина: Платёж успешно зафиксирован!")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go ProcessPayment(&wg)

	fmt.Println("Главный поток main: горутина запущена, включаю режим ожидания...")

	wg.Wait()

	fmt.Println("Главный поток main: Все потоки завершились, закрываю сервер!")
}
