package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Routes - all the registered routes
func Routes(router chi.Router) {
	router.Get("/", IndexHandeler)
	router.Get("/top", HealthHandeler)
	router.Route("/v1", InitV1Routes)
}

func InitV1Routes(r chi.Router) {
	r.Method(http.MethodGet, "/topics", Handler(getAllTopicsHandler))
	r.Method(http.MethodGet, "/topics/analysis", Handler(getTopicAnalysisHandler))
}
