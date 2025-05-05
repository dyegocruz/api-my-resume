package main

import "github.com/gofrs/uuid/v5"

type Position struct {
	ID        uuid.UUID `json:"_id"`
	Company   string    `json:"company"`
	Title     string    `json:"title"`
	Skills    []string  `json:"skills"`
	IsCurrent bool      `json:"isCurrent"`
	StartDate string    `json:"start_date"`
	EndDate   *string   `json:"end_date"` // Usando ponteiro para permitir valores nulos
	Summary   []string  `json:"summary"`
}

type Education struct {
	ID          uuid.UUID `json:"_id"`
	CourseName  string    `json:"course_name"`
	Degree      string    `json:"degree"`
	FromYear    string    `json:"from_year"`
	Institution string    `json:"institution"`
	ToYear      string    `json:"to_year"`
}

type Social struct {
	Github   string `json:"github"`
	Linkedin string `json:"linkedin"`
}

type MyResume struct {
	ID              uuid.UUID   `json:"_id"`
	Username        string      `json:"username"`
	Name            string      `json:"name"`
	Email           string      `json:"email"`
	Location        string      `json:"location"`
	Bio             string      `json:"bio"`
	AvatarURL       string      `json:"avatar_url"`
	CurrentPosition string      `json:"current_position"`
	Skills          []string    `json:"skills"`
	Social          Social      `json:"social"`
	Positions       []Position  `json:"positions"`
	Education       []Education `json:"education"`
}
