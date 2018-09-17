package domain

type User struct {
	Model
	Name        string `json:"name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}
