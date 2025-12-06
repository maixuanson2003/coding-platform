package entity

import "time"

type ListProblem struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"type:text"`
	Description string `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Problems []Problem `gorm:"many2many:list_problem_items;constraint:OnDelete:CASCADE;"`
}
