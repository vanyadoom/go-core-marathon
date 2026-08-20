package main

import "fmt"

func main() {
	totalBill := 4550.75
	friends := float64(4)
	share := totalBill / friends
	fmt.Printf("Каждый должен будет заплатить: %.2f", share)
}
