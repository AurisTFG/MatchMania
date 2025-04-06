package models

type Role string

const (
	AdminRole     Role = "admin"
	ModeratorRole Role = "moderator"
	UserRole      Role = "user"
)

func (r *Role) Scan(value any) error {
	*r = Role(value.(string))
	return nil
}

func (r *Role) Value() (any, error) {
	return string(*r), nil
}
