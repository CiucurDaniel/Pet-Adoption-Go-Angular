package entities

type User struct {
	ID uint `json:"id"`
	Email string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Posts []Post
}