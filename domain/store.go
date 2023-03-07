package domain

// Category Model using struct
type Category struct {
	CategoryID string `json:"categoryID"`
	Name       string `json:"name"`
}

func (Category) TableName() string {
	return "categories"
}

// Product represents a product entity.
type Product struct {
	ID        int64   `json:"id"`
	SKU       int64   `json:"sku"`
	Name      string  `json:"name"`
	SalePrice float64 `json:"sale_price"`
	Images    []struct {
		Href string `json:"href"`
	} `json:"images"`
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
	IsCategoryExist(category Category) (exist bool, err error)
	AddCategory(category Category) (err error)
}
