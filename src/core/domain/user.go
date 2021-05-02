package domain

type User struct {
	DisplayName string
	Handle      string
	Mail        string
}

type UserCreationRequest struct {
	User        User
	Credentials Credentials
}

type UserResponse struct {
	Status Status

	User *User
}
