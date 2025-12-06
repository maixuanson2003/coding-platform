package entity

import "time"

type Problem struct {
	Id           uint   `gorm:"primaryKey;autoIncrement"`
	Category     string `gorm:"type:varchar(255)"`
	Difficult    string `gorm:"type:varchar(50)"`
	Title        string `gorm:"type:text"`
	Content      string `gorm:"type:text"`
	IsDeleted    bool   `gorm:"default:false"`
	IsDailyToday bool   `gorm:"default:false"`
	PointDaily   int    `gorm:"default:0"`
	MemoryLimit  int
	TimeLimit    int
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Testcases    []TestCase
	Submissions  []Submission
	ListProblems []ListProblem `gorm:"many2many:list_problem_items"`
}
