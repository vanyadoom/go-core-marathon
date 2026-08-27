package main

import "fmt"

func main() {

	storage := []any{"Etherium", 2026, 3.14}

	for _, data := range storage {

		fmt.Println(data)

	}
}
