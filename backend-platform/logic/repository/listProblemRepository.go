package repository

import (
	"lietcode/logic/entity"

	"gorm.io/gorm"
)

type ListProblemRepository struct {
	*Repository[entity.ListProblem]
}

func NewListProblemRepository(db *gorm.DB) *ListProblemRepository {
	return &ListProblemRepository{
		Repository: &Repository[entity.ListProblem]{DataAccess: db},
	}
}
