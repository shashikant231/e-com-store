package usecase

import (
	"e-commerce-store/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// StoreUseCase is the usecase for Store
type StoreUseCase struct {
	storeRepo domain.StoreRepository
}

// SyncCategory ...
func (s *StoreUseCase) SyncCategory(limit string, page string) (err error) {
	url := "https://stageapi.monkcommerce.app/task/categories?limit=" + limit + "&" + "page=" + page
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("x-api-key", "s72rash8762s31")
	// Use the http.Client to send the request and retrieve the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&domain.CategoriesResponse)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}
	// var categories []domain.Category
	for _, category := range domain.CategoriesResponse.Categories {
		fmt.Println(category.Name)
		exist, err := s.storeRepo.IsCategoryExist(category.ID)
		if err != nil {
			return err
		}
		if !exist {
			// add category to Category Table
			err = s.storeRepo.AddCategory(category)
		}
	}
	return err
}

// SyncCategory ...
func (s *StoreUseCase) SyncProduct(limit string, page string, id string) (err error) {
	url := "https://stageapi.monkcommerce.app/task/products?limit=" + limit + "&" + "page=" + page + "&" + "categoryID=" + id
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("x-api-key", "s72rash8762s31")
	// Use the http.Client to send the request and retrieve the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	// bodyString := string(bodyBytes)
	// fmt.Println("API Response as String:\n" + bodyString)

	err = json.Unmarshal(bodyBytes, &domain.ProductsResponse)
	if err != nil {
		fmt.Println("eRROR WHILE DECODING")
	}
	// fmt.Println(&domain.ProductsResponse)
	for _, product := range domain.ProductsResponse.Products {
		_, err := s.storeRepo.IsProductExist(product.SKU)
		if err != nil {
			return err
		}
	}
	var products []domain.Product
	for _, product := range domain.ProductsResponse.Products {
		products = append(products, domain.Product{
			SKU:                 product.SKU,
			Name:                product.Name,
			SalePrice:           product.SalePrice,
			Digital:             product.Digital,
			CategoryID:          id,
			ShippingCost:        product.ShippingCost,
			Description:         product.Description,
			CustomerReviewCount: product.CustomerReviewCount,
		})
	}
	err = s.storeRepo.AddProduct(products)
	if err != nil {
		return
	}
	return err
}

// NewStoreUseCase creates will create new an Storeusecase object representation of domain.StoreUsecase interface
func NewStoreUseCase(StoreRepo domain.StoreRepository) domain.StoreUseCase {
	return &StoreUseCase{
		storeRepo: StoreRepo,
	}
}
