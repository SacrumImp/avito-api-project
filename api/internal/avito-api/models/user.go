package models

type UserRegisterObject struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}

type UserLogin struct {
	UserId string `json:"user_id"`
}

type UserAccount struct {
	UserId       string
	Email        string
	PasswordHash string
	UserType     string
}
