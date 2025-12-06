package repository

import (
	"lietcode/logic/entity"

	"gorm.io/gorm"
)

type SubmissionRepository struct {
	*Repository[entity.Submission]
}

func NewSubmissionRepository(db *gorm.DB) *SubmissionRepository {
	return &SubmissionRepository{
		Repository: &Repository[entity.Submission]{DataAccess: db},
	}
}


