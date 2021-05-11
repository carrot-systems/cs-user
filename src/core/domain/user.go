package domain

type User struct {
	ID          string
	DisplayName string `json:"display_name" binding:"required"`
	Handle      string `json:"handle" binding:"required"`
	Mail        string `json:"mail" binding:"required"`
}

type UserCreationRequest struct {
	User        User        `json:"user" binding:"required"`
	Credentials Credentials `json:"credentials" binding:"required"`
}

type UserResponse struct {
	Status Status

	User *User
}
