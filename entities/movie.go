package entities

import "time"

type Movie struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Duration   string    `json:"duration"`
	RelaseDate time.Time `json:"release_date"`
}
