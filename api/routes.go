package api

import "github.com/go-chi/chi"

// Routes - all the registered routes
func Routes(router chi.Router) {
	router.Get("/", IndexHandeler)
	router.Get("/top", HealthHandeler)
	// router.Route("/v1", Init)
}
