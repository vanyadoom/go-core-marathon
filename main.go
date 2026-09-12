package main

import (
	"context"
	"fmt"
)

type congigKey string

const reqIDkey congigKey = "x-request-id"

func GetMovieData(ctx context.Context, title string) {
	rawID := ctx.Value(reqIDkey)
	idStr, ok := rawID.(string)
	if !ok {
		idStr = "UNKNOWN_ID"
	}
	fmt.Printf("[ID: %s] Успешно прочитан фильм: %s\n", idStr, title)
}

func main() {
	ctx := context.WithValue(context.Background(), reqIDkey, "REQ-999-ONLINE")
	GetMovieData(ctx, "Интерстеллар")
}
