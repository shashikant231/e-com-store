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
	// url := "https://stageapi.monkcommerce.app/task/categories?limit=" + "100" + "&" + "page=" + %pages
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
		err = json.NewDecoder(resp.Body).Decode(&domain.CategoriesResponse)
		if err != nil {
			fmt.Println("Error decoding response:", err)
			return err
		}
		// var categories []domain.Category
		for _, category := range domain.CategoriesResponse.Categories {
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
			if len(domain.ProductsResponse.Products) == 0 {
				continue
			}

			err = json.Unmarshal(bodyBytes, &domain.ProductsResponse)
			if err != nil {
				fmt.Println("eRROR WHILE DECODING")
			}

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
		if len(domain.CategoriesResponse.Categories) == 0 {
			return err
		}
		pages++
	}
}

// NewStoreUseCase creates will create new an Storeusecase object representation of domain.StoreUsecase interface
func NewStoreUseCase(StoreRepo domain.StoreRepository) domain.StoreUseCase {
	return &StoreUseCase{
		storeRepo: StoreRepo,
	}
}
