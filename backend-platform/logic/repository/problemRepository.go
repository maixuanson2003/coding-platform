package repository

import (
	"lietcode/logic/entity"

	"gorm.io/gorm"
)

type ProblemRepository struct {
	*Repository[entity.Problem]
}

func NewProblemRepository(db *gorm.DB) *ProblemRepository {
	return &ProblemRepository{
		Repository: &Repository[entity.Problem]{DataAccess: db},
	}
}
func (r *ProblemRepository) GetListProblem(category *string, difficult *string, title *string, preLoad []string) ([]entity.Problem, error) {
	db := r.DataAccess.Model(&entity.Problem{})

	if category != nil && *category != "" {
		db = db.Where("category = ?", *category)
	}

	if difficult != nil && *difficult != "" {
		db = db.Where("difficult = ?", *difficult)
	}

	if title != nil && *title != "" {
		db = db.Where("title LIKE ?", "%"+*title+"%")
	}

	for _, preload := range preLoad {
		db = db.Preload(preload)
	}

	var problems []entity.Problem
	err := db.Find(&problems).Error
	return problems, err
}
