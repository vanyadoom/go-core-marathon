package main

import "fmt"

func main() {
	a := 5.0
	b := 5.0
	c := 8.0

	if a+b > c && a+c > b && b+c > a {

		if a == b && b == c {
			fmt.Println("Тип: Равносторонний треугольник")
		} else if a == b || b == c || a == c {
			fmt.Println("Тип: Равнобедренный треугольник")
		} else {
			fmt.Println("Тип: Разносторонний треугольник")
		}

	} else {

		fmt.Println("Ошибка: Треугольник с такими сторонами не существует!")
	}
}
