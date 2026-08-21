package main

import "fmt"

func main() {
	client1 := CalculateDelivery(2.5)
	client2 := CalculateDelivery(10.0)
	client3 := CalculateDelivery(0.5)

	fmt.Printf("Первый клиент: %.2f\n", client1)
	fmt.Printf("Второй клиент: %.2f\n", client2)
	fmt.Printf("Третий клиент: %.2f\n", client3)
}

func CalculateDelivery(weight float64) float64 {

	BasePrice := 300.0 + (weight * 50.5)

	return BasePrice
}
