package auth

type LoginDto struct {
	Email    string `example:"email@example.com"  json:"email"    validate:"required,email"`
	Password string `example:"VeryStrongPassword" json:"password" validate:"required"`
}
