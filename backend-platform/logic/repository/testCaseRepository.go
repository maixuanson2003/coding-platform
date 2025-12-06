package repository

import (
	"lietcode/logic/entity"

	"gorm.io/gorm"
)

type TestcaseRepository struct {
	*Repository[entity.TestCase]
}

func NewTestcaseRepository(db *gorm.DB) *TestcaseRepository {
	return &TestcaseRepository{
		Repository: &Repository[entity.TestCase]{DataAccess: db},
	}
}
