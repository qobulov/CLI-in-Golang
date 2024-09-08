package models

type Tasks struct {
	Day    int    `json:"day" db:"day"`
	Works  string `json:"works" db:"workout_description"`
	IsDone bool   `json:"isDone" db:"done"`
}
