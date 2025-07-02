package models

import (
	"github.com/google/uuid"
)

type Comment struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	MovieID uuid.UUID `gorm:"type:uuid;not null" json:"movieId"`
	UserID  uuid.UUID `gorm:"type:uuid;not null" json:"userId"`
	Content string    `gorm:"type:text;not null" json:"content"`
}
