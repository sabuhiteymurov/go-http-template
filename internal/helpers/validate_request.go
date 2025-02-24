package helpers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func HandleValidationError(w http.ResponseWriter, err error) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, validationErr := range validationErrors {
			field := validationErr.StructField()
			errors[field] = validationErr.Tag()
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Validation failed",
			"errors":  errors,
		})
		return
	}

	http.Error(w, "Internal server error", http.StatusInternalServerError)
}

func ValidateRequestData(w http.ResponseWriter, data interface{}) bool {
	customValidator := NewCustomValidator()
	if err := customValidator.validator.Struct(data); err != nil {
		HandleValidationError(w, err)
		return false
	}
	return true
}
