package dto

type RegisterEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type RegisterContactRequest struct {
	Contact string `json:"contact" validate:"required"`
}

type AddInformationRequest struct {
	TempToken string `json:"temp_token" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname"`
}

type AddPasswordRequest struct {
	TempToken       string `json:"temp_token" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type VerifyEmailRequest struct {
	TempToken string `json:"temp_token" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}
