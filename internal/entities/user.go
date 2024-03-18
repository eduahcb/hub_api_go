package entities

import "time"

type User struct {
	ID        uint
	Name      string `gorm:"not null; size:50"`
	Email     string `gorm:"not null;unique"`
	Password  string `gorm:"not null"`
	Contact   string
	Bio       string
  ModuleID  uint
  Module    Module
	CreatedAt time.Time
	UpdatedAt time.Time
}
