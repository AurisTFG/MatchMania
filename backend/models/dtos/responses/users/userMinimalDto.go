package responses

import (
	"github.com/google/uuid"
)

type UserMinimalDto struct {
	Id       uuid.UUID `example:"526432ea-822b-4b5b-b1b3-34e8ab453e03" json:"id"`
	Username string    `example:"john_doe_123"                         json:"username"`
}
