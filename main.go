package main

import (
	"fmt"
)

type Asset struct {
	Ticker   string
	ValueUSD float64
}

func main() {

	portfolio := []Asset{
		{Ticker: "BTC", ValueUSD: 1200.50},
		{Ticker: "ETH", ValueUSD: 450.00},
		{Ticker: "SOL", ValueUSD: 150.25},
	}
	totalPortfolioSum := 0.0

	for _, item := range portfolio {
		totalPortfolioSum += item.ValueUSD
	}

	fmt.Printf("Стоимость портфеля: %.2f\n", totalPortfolioSum)

}
