package entity

import "time"

type CodeTemplate struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Content   string `gorm:"type:text"`
	Lang      string `gorm:"type:varchar(255)"` // CSV hoặc JSON
	CreatedAt time.Time
	UpdatedAt time.Time
}
