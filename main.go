package main

import (
	"context"
	"fmt"
	"time"
)

func DownloadMovie(ctx context.Context, title string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("🛑 Скачивание фильма '%s' прервано: %v\n", title, ctx.Err())
			return

		default:
			fmt.Println("Качаю следующий гигабайт...")
			time.Sleep(20 * time.Millisecond)
		}
	}
}

func main() {
	ctxBackground := context.Background()
	ctx, cancel := context.WithCancel(ctxBackground)
	go DownloadMovie(ctx, "Интерстеллар")
	time.Sleep(50 * time.Millisecond)
	cancel()
	time.Sleep(50 * time.Millisecond)
}
