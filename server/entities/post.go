package entities

type Post struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	County      string `json:"county"`
	City        string `json:"city"`
	PhoneNumber string `json:"phone_number"`
	UserID uint `json:"user_id"`
}
