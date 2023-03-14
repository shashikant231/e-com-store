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

// Sync ...
func (s *StoreUseCase) Sync() (err error) {
	url := "https://stageapi.monkcommerce.app/task/categories?limit=100&page=%d"
	pages := 1
	for {
		req, err := http.NewRequest("GET", fmt.Sprintf(url, pages), nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return err
		}
		req.Header.Set("x-api-key", "s72rash8762s31")
		// Use the http.Client to send the request and retrieve the response
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return err
		}
		defer resp.Body.Close()
		var categoriesResponse domain.CategoriesResponse
		err = json.NewDecoder(resp.Body).Decode(&categoriesResponse)
		if err != nil {
			fmt.Println("Error decoding response:", err)
			return err
		}
		for _, category := range categoriesResponse.Categories {
			exist, err := s.storeRepo.IsCategoryExist(category.ID)
			if err != nil {
				return err
			}
			if !exist {
				// add category to Category Table
				err = s.storeRepo.AddCategory(category)
			}
			productUrl := fmt.Sprintf("https://stageapi.monkcommerce.app/task/products?limit=100&page=%d&categoryID=%s", 1, category.ID)
			req, err := http.NewRequest("GET", productUrl, nil)
			if err != nil {
				fmt.Println("Error creating request:", err)
				return err
			}
			req.Header.Set("x-api-key", "s72rash8762s31")
			// Use the http.Client to send the request and retrieve the response
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Error sending request:", err)
				return err
			}
			defer resp.Body.Close()
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			var productsRequest domain.ProductsRequest
			err = json.Unmarshal(bodyBytes, &productsRequest)
			if err != nil {
				fmt.Println("eRROR WHILE DECODING")
				return err
			}
			if len(productsRequest.Products) == 0 {
				continue
			}

			for _, product := range productsRequest.Products {
				_, err := s.storeRepo.IsProductExist(product.SKU)
				if err != nil {
					return err
				}
			}
			var products []domain.Product
			for _, product := range productsRequest.Products {
				products = append(products, domain.Product{
					SKU:                 product.SKU,
					Name:                product.Name,
					SalePrice:           product.SalePrice,
					Digital:             product.Digital,
					CategoryID:          category.ID,
					ShippingCost:        product.ShippingCost,
					Description:         product.Description,
					CustomerReviewCount: product.CustomerReviewCount,
				})
			}
			err = s.storeRepo.AddProduct(products)
			if err != nil {
				return err
			}
		}
		if len(categoriesResponse.Categories) == 0 {
			return err
		}
		pages++
	}
}

// GetCategories retrieve existing categories in database
func (s *StoreUseCase) GetCategories(limit uint, page uint) (categoriesResponse domain.CategoriesResponse, err error) {
	categories, err := s.storeRepo.GetCategories(limit, page)
	categoriesResponse = domain.CategoriesResponse{
		Page:       int(page),
		Categories: categories,
	}

	return
}

// GetProducts retrieve existing Products in database
func (s *StoreUseCase) GetProducts(limit uint, page uint, categoryID string) (productsResponse domain.ProductsResponse, err error) {
	products, err := s.storeRepo.GetProducts(limit, page, categoryID)
	productsResponse = domain.ProductsResponse{
		Page:     int(page),
		Products: products,
	}

	return
}

// NewStoreUseCase creates will create new an Storeusecase object representation of domain.StoreUsecase interface
func NewStoreUseCase(StoreRepo domain.StoreRepository) domain.StoreUseCase {
	return &StoreUseCase{
		storeRepo: StoreRepo,
	}
}
