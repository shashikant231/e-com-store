package domain

// Category Model using struct
type Category struct {
	ID   string `json:"ID" gorm:"primarykey;column:id"`
	Name string `json:"name" gorm:"column:name"`
}

func (Category) TableName() string {
	return "categories"
}

type Image struct {
	ID         uint   `json:"id" gorm:"primarykey;column:id"`
	Href       string `json:"href"`
	ProductSKU int64  `json:"-" gorm:"column:product_sku"`
}

type ImageResonse struct {
	Href string `json:"href"`
}

// ProductResponse represents a response of product entity.
type ProductResponse struct {
	SKU                 int64          `json:"sku" gorm:"primarykey;column:sku"`
	Name                string         `json:"name" gorm:"column:name"`
	SalePrice           float64        `json:"salePrice" gorm:"column:salePrice"`
	Images              []ImageResonse `json:"images"`
	Digital             bool           `json:"digital"`
	ShippingCost        int64          `json:"shippingCost" gorm:"column:shippingCost"`
	Description         *string        `json:"description"`
	CustomerReviewCount *int           `json:"customerReviewCount" gorm:"column:customerReviewCount"`
}

// Product represents a product entity.
type Product struct {
	SKU       int64   `json:"sku" gorm:"primarykey;column:sku"`
	Name      string  `json:"name" gorm:"column:name"`
	SalePrice float64 `json:"salePrice" gorm:"column:salePrice"`
	// Images              []Image `json:"images" gorm:"foreignKey:ProductSKU"`
	Digital             bool    `json:"digital"`
	CategoryID          string  `json:"category_id" gorm:"column:category_id"`
	ShippingCost        int64   `json:"shippingCost" gorm:"column:shippingCost"`
	Description         *string `json:"description"`
	CustomerReviewCount *int    `json:"customerReviewCount" gorm:"column:customerReviewCount"`
}

func (Product) TableName() string {
	return "products"
}

// CategoriesResponse represent the response for Category
var CategoriesResponse struct {
	Page       int        `json:"page"`
	Categories []Category `json:"categories"`
}

var ProductsResponse struct {
	Page     int               `json:"page"`
	Products []ProductResponse `json:"products"`
}

// StoreUseCase interface - business process handeler
type StoreUseCase interface {
	SyncCategory(limit string, page string) (err error)
	SyncProduct(limit, page, id string) (err error)
}

// StoreRepository interface - Crud operation
type StoreRepository interface {
	IsCategoryExist(categoryID string) (exist bool, err error)
	AddCategory(category Category) (err error)
	IsProductExist(sku int64) (exist bool, err error)
	AddProduct(products []Product) (err error)
}
