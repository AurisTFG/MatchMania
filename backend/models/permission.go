package models

type Permission struct {
	BaseModel

	Name        string `gorm:"unique;not null"`
	Description string `gorm:"not null"`

	Roles []Role `gorm:"many2many:role_permissions;"`
}
