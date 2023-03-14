package domain

// ErrorDetails is a struct used for storing response of error details
type ErrorDetails struct {
	Code        string `json:"errorCode"`
	Description string `json:"errorDescription"`
}

func (e ErrorDetails) Error() string {
	return e.Code
}

// DuplicateProductError is used for storing error when product is already in database
var DuplicateProductError = ErrorDetails{Code: "duplicateProductError", Description: "Given product is already in database"}

// DuplicateCategoryError is used for storing error when Category is already in database
var DuplicateCategoryError = ErrorDetails{Code: "duplicateCategoryError", Description: "Given Category is already in database"}
