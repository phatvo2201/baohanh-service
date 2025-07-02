package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Username string    `gorm:"unique;not null" json:"username"`
	Phone    string    `gorm:"unique;not null" json:"phone"`
	Gender   string    `gorm:"default:other" json:"gender"`

	Email    string `gorm:"unique;not null" json:"email"`
	Password string `json:"password"`
	Role     string `gorm:"default:user" json:"role"` // Default role is "user"
}
