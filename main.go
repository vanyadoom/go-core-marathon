package main

import "fmt"

func main() {

	salary := 120000.50
	taxRate := 0.0

	if salary <= 50000.0 {
		taxRate = 0.0
	} else if salary <= 150000.0 {
		taxRate = 0.13
	} else {
		taxRate = 0.15
	}

	taxAmount := salary * taxRate

	netIncome := salary - taxAmount

	fmt.Printf("Грязная зарплата: %.2f .", salary)
	fmt.Printf("Сумма налога: %.2f .", taxAmount)
	fmt.Printf("Чистый доход: %.2f .", netIncome)

}
