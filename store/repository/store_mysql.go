package repository

import (
	"context"
	"e-commerce-store/domain"
	"errors"

	"gorm.io/gorm"
)

// mysqlStoreRepository
type mysqlStoreRepository struct {
	db *gorm.DB
}

// SyncCategory...
func (m *mysqlStoreRepository) IsCategoryExist(category domain.Category) (exist bool, err error) {
	var existingCategory domain.Category
	err = m.db.WithContext(context.Background()).
		Where("category_id = ?", category.CategoryID).
		First(&existingCategory).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	return true, nil
}

// AddCategory
func (m *mysqlStoreRepository) AddCategory(category domain.Category) (err error) {
	err = m.db.Create(&category).Error
	if err != nil {
		return err
	}
	return
}

// NewMysqlStoreRepository creates an object that represents the store.repository interface
func NewMysqlStoreRepository(db *gorm.DB) domain.StoreRepository {
	return &mysqlStoreRepository{
		db: db,
	}
}
