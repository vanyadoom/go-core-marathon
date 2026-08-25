package main

import "fmt"

type Wallet struct {
	Owner   string
	Balance float64
}

func (w *Wallet) Deposit(amount float64) {
	w.Balance += amount
}

func main() {
	myWallet := Wallet{
		Owner:   "Ваня",
		Balance: 100.50,
	}

	fmt.Printf("Стартовый баланс: %.2f\n", myWallet.Balance)

	myWallet.Deposit(50.25)

	fmt.Printf("Финальный баланс: %.2f\n", myWallet.Balance)
}
