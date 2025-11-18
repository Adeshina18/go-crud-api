package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name     string    `json:"name" binding:"required"`
	Email    string    `gorm:"unique" json:"email" binding:"required"`
	Password string    `json:"password" binding:"required"`
}
