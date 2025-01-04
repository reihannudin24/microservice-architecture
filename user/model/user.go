package model

type User struct {
	Id           int64  `json:"id"`
	Email        string `json:"email"`
	Contact      string `json:"contact"`
	Username     string `json:"username"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Status       string `json:"status"`
	PasswordHash string `json:"password_hash"`
	EmailVerify  string `json:"email_verify"`
	IsVerify     bool   `json:"is_verify"`
	CreatedAt    string `json:"created_at"`
	UpdateAt     string `json:"update_at"`
}
