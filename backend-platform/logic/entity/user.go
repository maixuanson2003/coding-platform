package entity

import "time"

type User struct {
	Id           uint   `gorm:"primaryKey;autoIncrement"`
	Username     string `gorm:"type:varchar(25);not null"`
	Password     string `gorm:"type:varchar(100);not null"`
	Email        string `gorm:"type:varchar(100);not null"`
	Avatar       string `gorm:"type:varchar(255)"`
	NumberHandle int    `gorm:"default:0"`
	PointDaily   int    `gorm:"default:0"`
	IsActive     bool   `gorm:"default:true"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Submissions []Submission
}
