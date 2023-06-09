package model

import "time"

type Note struct {
	ID            int       `json:"id" db:"id"`
	UserId        int       `json:"user_id" db:"user_id"`
	Title         string    `json:"title" db:"title"`
	Content       string    `json:"content" db:"content"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	IsPublic      bool      `json:"is_public" db:"is_public"`
	LikeCount     int       `json:"like_count" db:"like_count"`
	FavoriteCount int       `json:"favorite_count" db:"favorite_count"`
}
