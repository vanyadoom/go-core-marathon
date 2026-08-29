package main

import (
	"fmt"
)

func main() {

	defer func() {
		// 🟢 Пытаемся поймать панику
		if r := recover(); r != nil {
			fmt.Println("Авария успешно ликвидирована. Причина паники:", r)
		}
	}()

	fmt.Println("Сервер запущен. Проверка безопасности...")

	panic("Обнаружена хакерская атака на ядро!")

}
