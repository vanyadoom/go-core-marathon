package main

import "fmt"

func main() {
	age := 16
	height := 155
	hasCard := false
	passesSafety := (age >= 14) && (height >= 155)
	canRide := passesSafety || hasCard
	fmt.Println("Допуск на аттракцион разрешен: ", canRide)
}
