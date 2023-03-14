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
func (m *mysqlStoreRepository) IsCategoryExist(categoryID string) (exist bool, err error) {
	var existingCategory domain.Category
	err = m.db.WithContext(context.Background()).
		Where("id = ?", categoryID).
		First(&existingCategory).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, domain.DuplicateCategoryError
}

// AddCategory
func (m *mysqlStoreRepository) AddCategory(category domain.Category) (err error) {
	err = m.db.Create(&category).Error
	if err != nil {
		return err
	}
	return
}

// SyncProduct...
func (m *mysqlStoreRepository) IsProductExist(sku int64) (exist bool, err error) {
	var existingProduct domain.Product
	err = m.db.WithContext(context.Background()).
		Where("sku = ?", sku).
		First(&existingProduct).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, domain.DuplicateProductError
}

// AddProduct
func (m *mysqlStoreRepository) AddProduct(products []domain.Product) (err error) {
	err = m.db.Create(&products).Error
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
