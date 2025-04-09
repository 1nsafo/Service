package main

import (
	"Service/internal/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	byteValue, err := os.ReadFile("data/movies_dataset.json")
	if err != nil {
		log.Fatalf("Ошибка чтения файла: %v", err)
	}

	var films []models.Film
	err = json.Unmarshal(byteValue, &films)
	if err != nil {
		log.Fatalf("Ошибка разбора JSON: %v", err)
	}

	for _, film := range films {
		fmt.Printf("ID: %d, Title: %s, ReleaseYear: %s, Director: %s, Genre: %s, Rating: %1f\n",
			film.ID, film.Title, film.ReleaseYear, film.Director, film.Genre, film.Rating)

	}
}
