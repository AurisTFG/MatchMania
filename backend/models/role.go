package models

type Role struct {
	BaseModel

	Name        string `gorm:"unique,not null"`
	Description string `gorm:"not null"`

	Permissions []Permission `gorm:"many2many:role_permissions;"`
	Users       []User       `gorm:"many2many:user_roles;"`
}
