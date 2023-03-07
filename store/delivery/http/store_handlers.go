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

	e.GET("/sync", handler.SyncCategory)
	// e.GET("/shop/categories",handler.GetCategories)
}

// SyncCategory to sync the catalog and product data.
func (s *StoreHandler) SyncCategory(c echo.Context) error {
	fmt.Println("request aya")
	limit := c.QueryParam("limit")
	page := c.QueryParam("page")
	err := s.StoreUsecase.SyncCategory(limit, page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, "Succesfully updated")
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
