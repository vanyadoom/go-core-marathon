package main

import (
	"encoding/json"
	"fmt"
)

type Asset struct {
	Ticker string  `json:"ticker"`
	Price  float64 `json:"price"`
}

func main() {

	rawJSON := []byte(`{"ticker":"SOL","price":145.20}`)

	var myAsset Asset

	fmt.Printf("До распаковки: Тикер: %s, Цена: %.2f\n", myAsset.Ticker, myAsset.Price)

	err := json.Unmarshal(rawJSON, &myAsset)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("После распаковки: Тикер: %s, Цена: %.2f\n", myAsset.Ticker, myAsset.Price)
	}

}
