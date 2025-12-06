package entity

import "time"

type Submission struct {
	Id        uint `gorm:"primaryKey;autoIncrement"`
	UserId    uint `gorm:"not null"`
	ProblemId uint `gorm:"not null"`

	Lang      string `gorm:"type:varchar(20)"`
	Code      string `gorm:"type:text"`
	Status    string `gorm:"type:varchar(20)"`
	RuntimeMS int
	MemoryKB  int
	CreatedAt time.Time

	User    User
	Problem Problem
}
