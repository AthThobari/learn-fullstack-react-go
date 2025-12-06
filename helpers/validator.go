package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func TranslateErrorMessage(err error) map[string]string {

	// Create a map to hold error messages
	errorsMap := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			field := fieldError.Field() //stores the names of fields that failed validation
			switch fieldError.Tag() {   //handle different types of validation
			case "required":
				errorsMap[field] = fmt.Sprintf("%s is required", field)
			case "email":
				errorsMap[field] = "Invalid email format"
			case "unique":
				errorsMap[field] = fmt.Sprintf("%s already exists", field)
			case "min":
				errorsMap[field] = fmt.Sprintf("%s must be at least %s characters", field, fieldError.Param())
			case "max":
				errorsMap[field] = fmt.Sprintf("%s must be at most %s characters", field, fieldError.Param())
			case "numeric":
				errorsMap[field] = fmt.Sprintf("%s must be a number", field)
			default:
				errorsMap[field] = "Invalid value"
			}
		}
	}

	// Handle error from GORM for duplicate entry
	if err != nil {
		// Check if the error contains "Duplicate entry" (duplicate data in the database)
		if strings.Contains(err.Error(), "Duplicate entry") {
			if strings.Contains(err.Error(), "username") {
				errorsMap["Username"] = "Username already exists"
			}
			if strings.Contains(err.Error(), "email") {
				errorsMap["Email"] = "Email already exists"
			}
		} else if err == gorm.ErrRecordNotFound {
			// If the data being searched for is not found in the database
			errorsMap["error"] = "Record not found"
		}
	}

	// Returns a map containing error messages
	return errorsMap
}

// IsDuplicateEntryError detects whether the error from the database is a duplicate entry
func IsDuplicateEntryError(err error) bool {
	// Checks whether the error is a duplicate entry
	return err != nil && strings.Contains(err.Error(), "Duplicate entry")
}
