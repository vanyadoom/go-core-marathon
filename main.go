package main

import "fmt"

func main() {
	rates := [5]float64{10.5, 45.2, 89.1, 23.7, 56.4}

	maxRate := rates[0]

	for i := 1; i < 5; i++ {
		if rates[i] > maxRate {
			maxRate = rates[i]
		}
	}
	fmt.Printf("%.1f\n", maxRate)
}
