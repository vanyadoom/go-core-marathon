package main

import (
	"fmt"
	"os"
)

func main() {

	os.Setenv("DB_PASSWORD", "aurora_secure_pass_2026")

	password := os.Getenv("DB_PASSWORD")

	if password == "" {
		fmt.Println("Критическая ошибка: Пароль базы данных не найден в системе!")
	} else {
		fmt.Printf("Успешное подключение к БД! Используется секретный пароль: %s\n", password)
	}

}
