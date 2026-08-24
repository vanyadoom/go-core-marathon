package main

import (
	"fmt"
)

type User struct {
	Name  string
	Coins int
}

func (u User) Showprofile() {
	fmt.Printf("Пользователь: %s, Баланс: %d BTC\n", u.Name, u.Coins)
}

func main() {
	vanya := User{Name: "Ваня", Coins: 3}
	vanya.Showprofile()
}
