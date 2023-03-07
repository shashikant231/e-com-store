package domain

// Category Model using struct
type Category struct {
	ID   string `json:"ID" gorm:"primarykey;column:id"`
	Name string `json:"name" gorm:"column:name"`
}

func (Category) TableName() string {
	return "categories"
}

// Image struct
type Image struct {
	Href string `json:"href"`
}

// Product represents a product entity.
type Product struct {
	ID                  int64   `json:"id"`
	SKU                 int64   `json:"sku"`
	Name                string  `json:"name"`
	SalePrice           float64 `json:"sale_price"`
	Images              []Image `json:"images" gorm:"-"`
	CategoryID          uint    `json:"category_id"`
	Digital             bool    `json:"digital"`
	ShippingCost        float64 `json:"shippingCost"`
	Description         string  `json:"description"`
	CustomerReviewCount int64   `json:"customerReviewCount"`
}

// CategoriesResponse represent the response for Category
var CategoriesResponse struct {
	Page       int        `json:"page"`
	Categories []Category `json:"categories"`
}

// StoreUseCase interface - business process handeler
type StoreUseCase interface {
	SyncCategory(limit string, page string) (err error)
}

// StoreRepository interface - Crud operation
type StoreRepository interface {
	IsCategoryExist(categoryID string) (exist bool, err error)
	AddCategory(category Category) (err error)
}
