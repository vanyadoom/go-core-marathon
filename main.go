package main

import "fmt"

func main() {

	transactions := []float64{50.0, 1200.50, 300.0, 4500.0, 80.25}

	whaleSum := 0.0

	for _, amount := range transactions {
		if amount > 500.0 {
			whaleSum += amount
		}
	}

	fmt.Printf("Китовые транзакции: %.2f\n", whaleSum)
}
