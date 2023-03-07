package usecase

import (
	"e-commerce-store/domain"
	"encoding/json"
	"net/http"
)

// StoreUseCase is the usecase for Store
type StoreUseCase struct {
	storeRepo domain.StoreRepository
}

// SyncCategory ...
func (s *StoreUseCase) SyncCategory(limit string, page string) (err error) {
	resp, err := http.Get("https://stageapi.monkcommerce.app/task/categories?limit=" + limit + "&" + "page=" + page)
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(&domain.CategoriesResponse)
	if err != nil {
		return err
	}
	// var categories []domain.Category
	for _, category := range domain.CategoriesResponse.Categories {
		exist, err := s.storeRepo.IsCategoryExist(category)
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
