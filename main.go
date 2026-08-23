package main

import "fmt"

func main() {
	var cart []float64

	cart = append(cart, 150.50, 300.00, 450.25)

	totalPrice := 0.0

	for i := 0; i < len(cart); i++ {
		totalPrice += cart[i]
	}

	fmt.Printf("Итоговая сумма корзины: %.2f\n", totalPrice)

}
