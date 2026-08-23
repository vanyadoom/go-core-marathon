package main

import "fmt"

func main() {

	balances := []float64{100.0, 520.50, 50.0, 1000.0}

	for index := range balances {
		balances[index] *= 1.5
	}

	for _, val := range balances {
		fmt.Printf("Новый баланс: %.2f\n", val)
	}
}
