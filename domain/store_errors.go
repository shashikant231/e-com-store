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
