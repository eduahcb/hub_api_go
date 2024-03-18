package entities

import "time"

type Tech struct {
	ID        uint
	Name      string `gorm:"not null; size:50"`
  LevelID   uint
  Level     Level
  UserID    uint
  User      User
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Tech) TableName() string {
    return "techs"
}

