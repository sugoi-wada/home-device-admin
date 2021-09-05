package model

import "time"

type CPDevice struct {
	ID        uint   `gorm:"primaryKey"`
	GatewayID string `gorm:"unique"`
	DeviceID  string `gorm:"unique"`
	Auth      string
	Nickname  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
