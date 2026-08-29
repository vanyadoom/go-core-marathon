package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	timestamp := time.Now().Unix()

	fmt.Println("Сгенерирован Unix-Timestamp:", timestamp)

	duration := time.Since(start)

	fmt.Printf("Скорость генерации лога: %s\n", duration)
}
