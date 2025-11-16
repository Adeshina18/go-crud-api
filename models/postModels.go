package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title     string    `json:"title" binding:"required"`
	Body      string    `json:"body"`
	CreatedAt time.Time
}
