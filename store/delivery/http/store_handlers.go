package http

import (
	"e-commerce-store/domain"
	"fmt"
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
	e.GET("/shop/categories", handler.GetCategories)
	e.GET("/shop/products", handler.GetProducts)
}

// Sync to sync the catalog and product data.
func (s *StoreHandler) Sync(c echo.Context) error {
	err := s.StoreUsecase.Sync()
	if err != nil {
		fmt.Println("ERR", err)
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	} else if err != nil && err == domain.DuplicateProductError {
		return c.JSON(http.StatusInternalServerError, domain.DuplicateProductError)
	} else if err != nil && err == domain.DuplicateCategoryError {
		return c.JSON(http.StatusInternalServerError, domain.DuplicateCategoryError)
	} else if err != nil && err == domain.ProductDecodingError {
		return c.JSON(http.StatusInternalServerError, domain.ProductDecodingError)
	} else if err != nil && err == domain.CategoryDecodingError {
		return c.JSON(http.StatusInternalServerError, domain.CategoryDecodingError)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Catalog data fetched and stored successfully",
	})
}

// GetCategories to retrieve the categories
func (s *StoreHandler) GetCategories(c echo.Context) error {
	limit, err := GetLimit(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	page, err := GetPage(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	categories, err := s.StoreUsecase.GetCategories(limit, page)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, categories)
}

// GetProducts to retrieve the Products
func (s *StoreHandler) GetProducts(c echo.Context) error {
	limit, err := GetLimit(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	page, err := GetPage(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	categoryID := c.QueryParam("categoryID")
	products, err := s.StoreUsecase.GetProducts(limit, page, categoryID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, products)
}
