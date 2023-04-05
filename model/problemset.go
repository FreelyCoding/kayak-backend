package model

import "time"

type ProblemSet struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	UserID      int       `json:"user_id" db:"user_id"`
	IsPublic    bool      `json:"is_public" db:"is_public"`
}
