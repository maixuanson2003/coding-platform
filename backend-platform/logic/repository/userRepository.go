package repository

import (
	"lietcode/logic/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	*Repository[entity.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Repository: &Repository[entity.User]{DataAccess: db},
	}
}
func (r *UserRepository) ExsistUserEmail(email string) (bool, error) {
	var count int64
	err := r.DataAccess.
		Model(&entity.User{}).
		Where("email = ?", email).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
