package models

import (
	"github.com/google/uuid"
)

type Rating struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	MovieID uuid.UUID `gorm:"type:uuid;not null" json:"movieId"`
	UserID  uuid.UUID `gorm:"type:uuid;not null" json:"userId"`
	Rating  float64   `gorm:"type:decimal(3,2);not null" json:"rating"`
}
