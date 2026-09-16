package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MovieRequest struct {
	Title    string  `json:"title"`
	Director string  `json:"director"`
	Rating   float64 `json:"rating"`
}

func AddMovieHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var movie MovieRequest
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, "Битый JSON-пакет", http.StatusBadRequest)
		return
	}

	// 🟢 ИСПРАВЛЕНО: Заменили несуществующий req на правильное имя movie
	fmt.Printf("🎬 Добавлен новый фильм: %s (Режиссёр: %s, Рейтинг: %.1f)\n", movie.Title, movie.Director, movie.Rating)

	// 🟢 ИСПРАВЛЕНО: Заменили Printf на Fprint(w), чтобы отправить ответ обратно клиенту по сети
	fmt.Fprint(w, "Фильм успешно принят бэкендом!")
}

func main() {
	fmt.Println("🚀 Юбилейный киносервер запущен на порту :8080...")

	http.HandleFunc("/addmovie", AddMovieHandler)
	http.ListenAndServe(":8080", nil)
}
