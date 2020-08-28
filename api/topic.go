package api

import (
	"aavaz/errors"
	"aavaz/respond"
	"net/http"
)

func getAllTopicsHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	topics, err := store.Topic().All()
	if err != nil {
		return err
	}

	respond.OK(w, topics)
	return nil
}
