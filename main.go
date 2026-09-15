package main

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Добро пожаловать в Go Core Кинотеатр!")
}

func MoviesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "1. Криминальное чтиво\n2. Бойцовский клуб\n3. Интерстеллар")
}

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/movies", MoviesHandler)
	fmt.Println("🚀 Киносервер запущен на порту :8080...")
	http.ListenAndServe(":8080", nil)
}
