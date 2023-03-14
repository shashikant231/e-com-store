package http

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetLimit a method to get deviceType from query params
func GetLimit(c echo.Context) (limit uint, err error) {

	idString := c.QueryParam("limit")
	uintID, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		return
	}
	limit = uint(uintID)

	return
}

// GetPage a method to get deviceType from query params
func GetPage(c echo.Context) (page uint, err error) {

	idString := c.QueryParam("page")
	uintID, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		return
	}
	page = uint(uintID)

	return
}
