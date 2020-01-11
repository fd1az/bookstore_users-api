package errors

import (
	"errors"
	"net/http"
)

//RestErr struct
type RestErr struct {
	Message string `json:"Message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewError(message string) error {
	return errors.New(message)
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
