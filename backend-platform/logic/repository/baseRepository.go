package repository

import (
	"lietcode/database"

	"gorm.io/gorm"
)

type Repository[T any] struct {
	DataAccess *gorm.DB
}

func NewRepository[T any]() *Repository[T] {
	return &Repository[T]{
		DataAccess: database.DatabaseInstance,
	}
}

// get list
func (repo *Repository[T]) FindAll(option map[string]interface{}, preload []string) ([]T, error) {
	database := repo.DataAccess
	var result []T
	var err error
	for _, item := range preload {
		database = database.Preload(item)
	}
	if option == nil {

		err = database.Find(&result).Error

	} else {
		err = database.Where(option).Find(&result).Error
	}
	return result, err
}

// find one
func (repo *Repository[T]) FindOne(option map[string]interface{}, preload []string) (T, error) {
	database := repo.DataAccess
	var result T
	for _, item := range preload {
		database = database.Preload(item)
	}
	err := database.Where(option).First(&result).Error
	return result, err
}

// create new record
func (repo *Repository[T]) Create(entity *T) (*T, error) {
	db := repo.DataAccess
	err := db.Create(entity).Error
	return entity, err
}

// Update updates the record with the given Id using the provided updatedData map.
func (repo *Repository[T]) Update(Id uint, updatedData map[string]interface{}) error {
	database := repo.DataAccess

	errs := database.Transaction(func(tx *gorm.DB) error {
		err := database.Model(new(T)).Where("id=?", Id).Updates(updatedData).Error
		if err != nil {
			return err
		}
		return nil
	})
	return errs
}

func (repo *Repository[T]) Delete(id uint) error {
	db := repo.DataAccess
	var entity T
	return db.Where("id = ?", id).Delete(&entity).Error
}
