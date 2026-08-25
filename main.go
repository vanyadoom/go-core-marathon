package main

import "fmt"

type Transaction struct {
	From   string
	Amount float64
}

func (t Transaction) String() string {
	return fmt.Sprintf("Перевод от %s: %.2f USD", t.From, t.Amount)
}

func main() {
	tx := Transaction{
		From:   "Алексей",
		Amount: 250.75,
	}

	fmt.Println(tx)

}
