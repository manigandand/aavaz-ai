package main

import (
	"aavaz/api"
	"aavaz/config"
	appmiddleware "aavaz/middleware"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

var (
	name    = "batman"
	version = "1.0.0"
)

func main() {
	api.InitAPI(name, version)

	router := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "OPTIONS", "DELETE"},
		AllowedHeaders: []string{
			"Origin", "Authorization", "Access-Control-Allow-Origin",
			"Access-Control-Allow-Header", "Accept",
			"Content-Type", "X-CSRF-Token",
		},
		ExposedHeaders: []string{
			"Content-Length", "Access-Control-Allow-Origin", "Origin",
		},
		AllowCredentials: true,
		MaxAge:           300,
	})
	// cross & loger middleware
	router.Use(cors.Handler)
	router.Use(
		middleware.Logger,
		appmiddleware.Recoverer,
	)

	router.Route("/", api.Routes)
	log.Infoln("Starting server on port:", config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.Port), router); err != nil {
		log.Fatal(err)
	}
}
