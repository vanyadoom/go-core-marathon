package main

import "fmt"

func main() {

	assets := make(map[string]float64)

	assets["BTC"] = 0.5
	assets["DOGE"] = 999.9

	delete(assets, "DOGE")

	_, ok := assets["DOGE"]

	if ok {
		fmt.Printf("Валюта DOGE до сих пор в портфеле!")
	} else {
		fmt.Printf("Успех: Валюта DOGE полностью удалена из портфеля!")
	}

}
