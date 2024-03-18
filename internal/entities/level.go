package entities

import "time"

type Level struct {
	ID        uint
	Name      string `gorm:"not null; size:50"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
