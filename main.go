package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Инструкция: Запустите программу с командой 'start' или 'stop'")
		return
	}

	command := os.Args[1]
	switch command {
	case "start":
		fmt.Println("Запуск бэкенд-системы Аврора...")
	case "stop":
		fmt.Println("Сервер успешно остановлен.")
	default:
		fmt.Println("Неизвестная команда!")
	}
}
