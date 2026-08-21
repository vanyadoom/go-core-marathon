package main

import "fmt"

func main() {

	q1 := 3.5
	q2 := 4.2
	q3 := 3.8

	average := (q1 + q2 + q3) / 3.0

	fmt.Printf("Средний балл: %.1f\n", average)

	if average >= 4.0 {

		fmt.Println("Результат: Отлично! Допуск к экзамену получен автоматические.")

	} else {

		fmt.Println("Результат: Внимание! Средний балл ниже 4.0. Требуется сдача зачета для допуска.")

	}
}
