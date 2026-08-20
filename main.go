package main

import "fmt"

func main() {
	pizzas := 5
	slicesPizza := 8
	totalSlices := pizzas * slicesPizza
	programmers := 6

	OnePieceOneProg := totalSlices / programmers

	fmt.Println("Каждому программисту достанется: ", OnePieceOneProg)

	ostatok := totalSlices % OnePieceOneProg

	fmt.Println("После того, как все покушают, в коробке останется: ", ostatok)
}
