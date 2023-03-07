package http

import (
	"e-commerce-store/domain"
	"net/http"

	// "e-commerce-store/store/delivery/http"

	"github.com/labstack/echo/v4"
)

// Store handler represent the httphandler for Store
type StoreHandler struct {
	StoreUsecase domain.StoreUseCase
}

// NewStoreHandler will initialize the Store endpoint
func NewStoreHandler(e *echo.Echo, us domain.StoreUseCase) {
	handler := &StoreHandler{
		StoreUsecase: us,
	}

	e.GET("/sync", handler.SyncCategory)
	// e.GET("/shop/categories",handler.GetCategories)
}

// SyncCategory to sync the catalog and product data.
func (s *StoreHandler) SyncCategory(c echo.Context) error {
	limit := c.QueryParam("limit")
	page := c.QueryParam("page")
	err := s.StoreUsecase.SyncCategory(limit, page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, "Succesfully updated")

	// resp, err := http.Get("https://stageapi.monkcommerce.app/task/categories")
	// if err != nil {
	// 	return err
	// }
	// err = json.NewDecoder(resp.Body).Decode(&categoriesResponse)
	// if err != nil {
	// 	return err
	// }
	// for _, category := range categoriesResponse.Categories {
	// 	// Check if the category already exists in the database
	// }
}

// func (s *StoreHandler) GetCategories(c echo.Context) error{
// 	limit := c.QueryParam("limit")
// 	page := c.QueryParam("page")
// 	err := s.StoreUsecase.GetCategories(limit, page)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, "Internal server error")
// 	}

// 	return c.JSON(http.StatusOK, "Succesfully updated")
// }
