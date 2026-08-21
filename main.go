package main

import "fmt"

func main() {
	totalSum := 0

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			totalSum += i

		}
	}

	fmt.Println("Сумма четных чисел от 1 до 10 равна: ", totalSum)

}
