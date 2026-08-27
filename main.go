package main

import "fmt"

func main() {

	var box any = 40

	value, ok := box.(int)

	if ok == true {
		result := value + 10
		fmt.Println("Успешная распаковка! Результат + 10 =", result)
	} else {
		fmt.Println("Ошибка: внутри коробки лежит не целое число!")
	}
}
