package main

import "fmt"

func main() {
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			if i == 3 || j == 3 {
				continue

			}
			fmt.Printf("%d x %d = %d\t", i, j, i*j)
		}
		if i != 3 {
			fmt.Println()
		}
	}
}
