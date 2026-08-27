package main

import "fmt"

func main() {

	var box any = "Solana"

	switch v := box.(type) {

	case int:
		fmt.Printf("Это целое число: %d\n", v)

	case string:
		fmt.Printf("Это текстовая строка: %s\n", v)

	default:
		fmt.Println("Неизвестный тип данных")

	}
}
