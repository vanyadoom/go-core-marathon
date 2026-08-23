package main

import "fmt"

func main() {

	portfolio := [][]string{

		{"Bitcoin", "2.5"},
		{"Ethereum", "10.0"},
		{"Solana", "150.75"},
	}

	for _, row := range portfolio {
		for _, cell := range row {
			fmt.Printf("%-12s", cell)
		}
		fmt.Println()
	}

}
