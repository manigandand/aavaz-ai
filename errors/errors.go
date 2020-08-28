package errors

import (
	"errors"
	"net/http"
)

// AppError struct holds the value of HTTP status code and custom error message.
type AppError struct {
	Status       int         `json:"status"`
	Message      string      `json:"error_message,omitempty"`
	Debug        error       `json:"-"`
	ConflictData interface{} `json:"conflict_data,omitempty"`
}

var ner = errors.New

// IsRequiredErr returns new error with custom error message
func IsRequiredErr(key string) error {
	return ner(key + " is required")
}

func (err *AppError) Error() string {
	return err.Message
}

// AddDebug method is used to add a debug error which will be printed
// during the error execution if it is not nil. This is purely for developers'
// debugging purposes
func (err *AppError) AddDebug(erx error) *AppError {
	if err != nil {
		err.Debug = erx
	}

	return err
}

// AddConflictData add extra conflict error information. This will be sent in meta
func (err *AppError) AddConflictData(data interface{}) *AppError {
	if err != nil {
		err.ConflictData = data
	}

	return err
}

// NewAppError returns the new apperror object
func NewAppError(status int, message string) *AppError {
	return &AppError{
		Status:  status,
		Message: message,
	}
}

// 4xx -------------------------------------------------------------------------

// BadRequest will return `http.StatusBadRequest` with custom message.
func BadRequest(message string) *AppError { // 400
	return NewAppError(http.StatusBadRequest, message)
}

// Unauthorized will return `http.StatusUnauthorized` with custom message.
func Unauthorized(message string) *AppError { // 401
	return NewAppError(http.StatusUnauthorized, message)
}

// Forbidden will return `http.StatusForbidden` with custom message.
func Forbidden(message string) *AppError { // 403
	return NewAppError(http.StatusForbidden, message)
}

// NotFound will return `http.StatusNotFound` with custom message.
func NotFound(message string) *AppError { // 404
	return NewAppError(http.StatusNotFound, message)
}

// Conflict will return `http.StatusConflict` with custom message.
func Conflict(message string) *AppError { // 409
	return NewAppError(http.StatusConflict, message)
}

// UnprocessableEntity will return `http.StatusUnprocessableEntity` with
// custom message.
func UnprocessableEntity(message string) *AppError { // 422
	return NewAppError(http.StatusUnprocessableEntity, message)
}

// TooManyRequests will return `http.StatusTooManyRequests` with
// custom message.
func TooManyRequests(message string) *AppError { // 422
	return NewAppError(http.StatusTooManyRequests, message)
}

// 5xx -------------------------------------------------------------------------

// InternalServer will return `http.StatusInternalServerError` with custom message.
func InternalServer(message string) *AppError { // 500
	return NewAppError(http.StatusInternalServerError, message)
}

// InternalServerStd will return `http.StatusInternalServerError` with static message.
func InternalServerStd() *AppError { // 500
	return NewAppError(http.StatusInternalServerError, "Something went wrong")
}

// IsStatusNotFound should return true if HTTP status of an error is 404.
func (err *AppError) IsStatusNotFound() bool {
	return err.Status == http.StatusNotFound
}

// IsBadRequest should return true if HTTP status of an error is 400.
func (err *AppError) IsBadRequest() bool {
	return err.Status == http.StatusBadRequest
}

// IsInternalServerError should return true if HTTP status of an error is 500.
func (err *AppError) IsInternalServerError() bool {
	return err.Status == http.StatusInternalServerError
}