package models

type UserRegisterObject struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}

type UserLogin struct {
	UserId int `json:"user_id"`
}

type UserAccount struct {
	UserId       int
	Email        string
	PasswordHash string
	UserType     string
}

type UserLoginObject struct {
	UserId   int    `json:"id"`
	Password string `json:"password"`
}
