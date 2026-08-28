package main

import (
	"errors"
	"fmt"
)

func main() {

	username := "Vanya"

	baseErr := errors.New("limit exceeded")

	advancedErr := fmt.Errorf("user %s financial error: %w", username, baseErr)

	if advancedErr != nil {
		fmt.Println(advancedErr)
	}

}
