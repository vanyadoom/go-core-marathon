package main

import "fmt"

type TaxCalculator interface {
	CalculateTax() float64
}

type Wallet struct {
	Balance float64
}

type CompanyAccount struct {
	Profit float64
}

func (w Wallet) CalculateTax() float64 {
	return w.Balance * 0.13
}

func (c CompanyAccount) CalculateTax() float64 {
	return c.Profit * 0.20
}

func main() {

	totalTax := 0.0

	assets := []TaxCalculator{
		Wallet{Balance: 2000.0},
		CompanyAccount{Profit: 5000.0},
	}

	for _, asset := range assets {
		totalTax += asset.CalculateTax() // 🟢 Automatically calls the correct method version!
	}

	fmt.Printf("%.2f\n", totalTax)
}
