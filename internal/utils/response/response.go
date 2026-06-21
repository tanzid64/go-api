package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "ERROR"
)

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		Status:  StatusError,
		Message: err.Error(),
		Success: false,
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errorMessages []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errorMessages = append(errorMessages, fmt.Sprintf("%s is required", err.Field()))
		case "email":
			errorMessages = append(errorMessages, fmt.Sprintf("%s must be a valid email", err.Field()))
		case "min":
			errorMessages = append(errorMessages, fmt.Sprintf("%s must be greater than or equal to %s", err.Field(), err.Param()))
		default:
			errorMessages = append(errorMessages, fmt.Sprintf("%s is invalid", err.Field()))
		}
	}
	return Response{
		Status:  StatusError,
		Message: "Validation failed",
		Success: false,
		Error:   strings.Join(errorMessages, ", "),
	}
}
