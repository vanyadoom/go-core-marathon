package main

import "fmt"

type Account struct {
	Owner   string
	Balance float64
}

func NewAccount(owner string, balance float64) Account {

	if balance < 0.0 {
		balance = 0.0
	}

	return Account{Owner: owner, Balance: balance}
}

func main() {

	myAcc := NewAccount("Дмитрий", -150.75)

	fmt.Printf("Владелец: %s, Баланс: %.2f\n", myAcc.Owner, myAcc.Balance)

}
