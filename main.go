package main

import "fmt"

type Person struct {
	Name string
}

type Trader struct {
	Person
	Balance float64
}

func main() {
	t := Trader{
		Person:  Person{Name: "Алексей"},
		Balance: 1500.50,
	}
	fmt.Printf("Трейлер: %s, Баланс: %.2f USD\n", t.Name, t.Balance)
}
