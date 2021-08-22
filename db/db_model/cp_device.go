package db_model

import "time"

type CPDevice struct {
	ID        uint `gorm:"primaryKey"`
	GatewayID string
	Auth      string
	DeviceID  string
	Nickname  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
