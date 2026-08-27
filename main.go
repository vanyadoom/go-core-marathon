package main

import (
	"errors"
	"fmt"
)

func Withdraw(amount float64) (string, error) {

	if amount <= 0.0 {
		return "", errors.New("сумма должна быть больше нуля")
	} else {
		return "Успешный вывод средств!", nil
	}
}

func main() {
	msg, err := Withdraw(-50.0)

	if err != nil {
		fmt.Println("Ошибка безопасности:", err)
	} else {
		fmt.Println(msg)
	}
}
