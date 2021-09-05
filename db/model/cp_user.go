package model

import "time"

type CPUser struct {
	ID           uint   `gorm:"primaryKey"`
	Email        string `gorm:"unique"`
	CPToken      string
	ExpireTime   string
	RefreshToken string
	MVersion     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
