package http

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

// Default represents default limits for different response types
var Default = map[string]string{
	"limit": "1",
	"page":  "1",
}

// GetLimit a method to get Limit from query params
func GetLimit(c echo.Context) (limit uint, err error) {

	limitString := c.QueryParam("limit")
	if len(limitString) == 0 {
		limitString = Default["limit"]
	}
	uintID, err := strconv.ParseUint(limitString, 10, 64)
	if err != nil {
		return
	}
	limit = uint(uintID)

	return
}

// GetPage a method to get Page from query params
func GetPage(c echo.Context) (page uint, err error) {

	pageString := c.QueryParam("page")
	if len(pageString) == 0 {
		pageString = Default["page"]
	}
	uintID, err := strconv.ParseUint(pageString, 10, 64)
	if err != nil {
		return
	}
	page = uint(uintID)

	return
}
