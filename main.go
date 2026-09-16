package main

import (
	"encoding/json"
	"net/http"
)

type Movie struct {
	Title    string  `json:"title"`
	Director string  `json:"director"`
	Rating   float64 `json:"rating"`
}

func MoviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	movies := []Movie{
		{Title: "Интерстеллар", Director: "Кристофер Нолан", Rating: 8.6},
		{Title: "Криминальное чтиво", Director: "Квентин Тарантино", Rating: 8.9},
		{Title: "Бойцовский клуб", Director: "Дэвид Финчер", Rating: 8.7},
	}

	json.NewEncoder(w).Encode(movies)
}

func main() {
	http.HandleFunc("/movies", MoviesHandler)
	http.ListenAndServe(":8080", nil)
}
