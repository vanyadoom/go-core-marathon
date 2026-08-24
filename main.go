package main

import (
	"fmt"
)

func main() {

	investments := make(map[string]float64)

	investments["Иван"] = 450.50
	investments["Алексей"] = 1200.00
	investments["Ольга"] = 350.25

	totalCapital := 0.0

	for _, balance := range investments {
		totalCapital += balance
	}

	fmt.Printf(" Итоговый капитал: %.2f\n", totalCapital)

}
