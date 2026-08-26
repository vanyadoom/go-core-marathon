package main

import "fmt"

type TaxCalculator interface {
	CalculateTax() float64 // Только сигнатура: имя, аргументы и что возвращает
}

type Wallet struct {
	Balance float64
}

func (w Wallet) CalculateTax() float64 {
	return w.Balance * 0.13
}

func main() {
	myWallet := Wallet{Balance: 1000.0}
	var calc TaxCalculator = myWallet

	fmt.Printf("Сумма налога: %.2f USD\n", calc.CalculateTax())
}
