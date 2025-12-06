package repository

import (
	"lietcode/logic/entity"

	"gorm.io/gorm"
)

type CodeTemplateRepository struct {
	*Repository[entity.CodeTemplate]
}

func NewCodeTemplateRepository(db *gorm.DB) *CodeTemplateRepository {
	return &CodeTemplateRepository{
		Repository: &Repository[entity.CodeTemplate]{DataAccess: db},
	}
}
