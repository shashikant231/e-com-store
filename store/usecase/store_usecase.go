package usecase

import (
	"e-commerce-store/domain"
	"encoding/json"
	"fmt"
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

// NewStoreUseCase creates will create new an Storeusecase object representation of domain.StoreUsecase interface
func NewStoreUseCase(StoreRepo domain.StoreRepository) domain.StoreUseCase {
	return &StoreUseCase{
		storeRepo: StoreRepo,
	}
}
