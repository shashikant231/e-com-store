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

// IsCategoryExist check if given Category exist in database or not
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

// AddCategory adds new product into database
func (m *mysqlStoreRepository) AddCategory(category domain.Category) (err error) {
	err = m.db.Create(&category).Error
	if err != nil {
		return err
	}
	return
}

// IsProductExist check if given product exist in database or not
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

// AddProduct adds new product into database
func (m *mysqlStoreRepository) AddProduct(products []domain.Product) (err error) {
	err = m.db.Create(&products).Error
	if err != nil {
		return err
	}
	return
}

// GetCategories retrieve existing categories in database
func (m *mysqlStoreRepository) GetCategories(limit uint, page uint) (categories []domain.Category, err error) {
	offset := (page - 1) * limit
	err = m.db.Table("categories").
		Select("categories.id, categories.name, COUNT(*) as product_count").
		Joins("LEFT JOIN products ON categories.id = products.category_id").
		Group("categories.id").
		Order("product_count DESC").
		Limit(int(limit)).
		Offset(int(offset)).
		Find(&categories).Error

	if err != nil {
		return
	}
	return

}

// GetProducts retrieve existing Products from database
func (m *mysqlStoreRepository) GetProducts(limit uint, page uint, categoryID string) (products []domain.Product, err error) {
	offset := (page - 1) * limit
	err = m.db.Order("customerReviewCount desc").
		Limit(int(limit)).
		Offset(int(offset)).
		Where("category_id = ?", categoryID).
		Find(&products).
		Error

	if err != nil {
		return
	}
	return

}

// NewMysqlStoreRepository creates an object that represents the store.repository interface
func NewMysqlStoreRepository(db *gorm.DB) domain.StoreRepository {
	return &mysqlStoreRepository{
		db: db,
	}
}
