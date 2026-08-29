package main

import (
	"encoding/json"
	"fmt"
)

type CryptoBalance struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

func main() {
	u := CryptoBalance{Currency: "BTC", Amount: 1.45}

	jsonData, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(jsonData))
	}
}
