package model

import "time"

type CPDeviceState struct {
	ID                  uint `gorm:"primaryKey"`
	CPDeviceID          uint
	Power               string
	Feature             string
	Speed               string
	Temp                string
	InsideTemp          string
	Nanoex              string
	People              string
	OutsideTemp         string
	PM25                string
	OnTimer             string
	OffTimer            string
	VerticalDirection   string
	HorizontalDirection string
	Fast                string
	Econavi             string
	Volume              string
	DisplayLight        string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
