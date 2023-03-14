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

	e.GET("/sync", handler.Sync)
	// e.GET("/shop/categories",handler.GetCategories)
}

// Sync to sync the catalog and product data.
func (s *StoreHandler) Sync(c echo.Context) error {
	err := s.StoreUsecase.Sync()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	} else if err != nil && err == domain.DuplicateProductError {
		return c.JSON(http.StatusInternalServerError, domain.DuplicateProductError)
	} else if err != nil && err == domain.DuplicateCategoryError {
		return c.JSON(http.StatusInternalServerError, domain.DuplicateCategoryError)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Catalog data fetched and stored successfully",
	})
}
