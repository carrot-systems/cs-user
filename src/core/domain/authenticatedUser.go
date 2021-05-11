package domain

import "github.com/google/uuid"

type AuthenticatedUser struct {
	ID uuid.UUID `json:"id" binding:"required"`
}

type AuthenticatedUserResponse struct {
	AuthenticatedUser `json:"authentication" binding:"required"`
}
