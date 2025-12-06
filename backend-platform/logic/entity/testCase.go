package entity

import (
	"time"
)

type TestCase struct {
	Id        uint `gorm:"primaryKey;autoIncrement"`
	ProblemId uint `gorm:"not null"`

	Input     string `gorm:"type:text"`
	Output    string `gorm:"type:text"`
	IsPublic  bool   `gorm:"default:false"`
	RuntimeMS int
	MemoryKB  int
	CreatedAt time.Time
	UpdatedAt time.Time

	Problem Problem
}
