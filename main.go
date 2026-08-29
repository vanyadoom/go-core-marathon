package main

import "fmt"

const (
	StatusPending = iota
	StatusProcessing
	StatusFailed
)

func main() {

	fmt.Println(StatusPending)

	fmt.Println(StatusProcessing)

	fmt.Println(StatusFailed)

}
