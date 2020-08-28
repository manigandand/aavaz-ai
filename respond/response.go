package respond

import (
	"aavaz/errors"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// SendResponse customize the http response and sends
func SendResponse(w http.ResponseWriter, status int, data interface{}) {
	res := &Response{
		Data: data,
		Meta: Meta{Status: status},
	}
	res.Send(w)
}

// 2xx -------------------------------------------------------------------------

// OK is a helper function used to send response data
// with StatusOK status code (200)
func OK(w http.ResponseWriter, data interface{}) {
	SendResponse(w, http.StatusOK, data)
}

// Created is a helper function used to send response data
// with StatusCreated status code (201)
func Created(w http.ResponseWriter, data interface{}) {
	SendResponse(w, http.StatusCreated, data)
}

// NoContent is a helper function used to send a NoContent Header (204)
// Note : the sent data and meta are ignored.
func NoContent(w http.ResponseWriter, data interface{}) {
	SendResponse(w, http.StatusNoContent, nil)
}

// 4xx & 5XX -------------------------------------------------------------------

// Fail write the error response
// Common func to send all the error response
func Fail(w http.ResponseWriter, e *errors.AppError) {
	log.Errorf("StatusCode: %d, Error: %s\n DEBUG: %s\n",
		e.Status, e.Error(), e.Debug)

	res := &Response{
		Data: nil,
		Meta: Meta{Status: e.Status, Message: e.Message},
	}
	res.Send(w)
}
