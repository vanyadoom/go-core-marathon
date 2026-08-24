package main

import "fmt"

func main() {

	cryptoWallet := make(map[string]float64)

	cryptoWallet["BTC"] = 0.45
	cryptoWallet["ETH"] = 3.20
	cryptoWallet["SOL"] = 45.75

	cryptoWallet["ETH"] += 1.5

	fmt.Printf("Баланс ETH: %.2f\n", cryptoWallet["ETH"])

}
