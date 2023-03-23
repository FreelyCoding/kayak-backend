package model

import "time"

type ProblemType struct {
	ID            int       `json:"id" db:"id"`
	Description   string    `json:"description" db:"description"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	UserID        int       `json:"user_id" db:"user_id"`
	ProblemTypeID int       `json:"problem_type_id" db:"problem_type_id"`
	IsPublic      bool      `json:"is_public" db:"is_public"`
}
type ProblemChoice struct {
	ID          int    `json:"id" db:"id"`
	Choice      string `json:"choice" db:"choice"`
	Description string `json:"description" db:"description"`
	IsCorrect   bool   `json:"is_correct" db:"is_correct"`
}
