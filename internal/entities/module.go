package entities

import "time"

type Module struct {
	ID        uint
  Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
