package dto

import "github.com/google/uuid"

type CreatePost struct {
	UserID   uuid.UUID `json:"user_id" validate:"required"`
	Content  string    `json:"content" validate:"required"`
	ImageURL string    `json:"image_url"`
}
