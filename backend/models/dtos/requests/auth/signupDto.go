package auth

type SignupDto struct {
	Username string `example:"john_doe_123"       json:"username" validate:"required,min=3,max=100"`
	Email    string `example:"email@example.com"  json:"email"    validate:"required,email"`
	Password string `example:"VeryStrongPassword" json:"password" validate:"required,min=8,max=100"`
}
