package main

import (
	"context"
	"fmt"
	"time"
)

func FetchMovieFromDB(ctx context.Context, title string) {
	time.Sleep(100 * time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Printf("❌ Запрос отменён фильтром безопасности: %v\n", ctx.Err())
		return
	default:
		fmt.Printf("✅ Успешно получили данные фильма: %s\n", title)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	FetchMovieFromDB(ctx, "Бэтмен")
	fmt.Println("Работа главного потока завершена.")
}
