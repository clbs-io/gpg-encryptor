package service

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2" // http-swagger middleware

	"github.com/cybroslabs/gpg-encryptor/internal/handlers"
)

func Service() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/healthz"))

	r.Mount("/swagger", swaggerRouter())
	r.Mount("/v1", apiRouter())

	return r
}

// return router for swagger docs
func swaggerRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The url pointing to API definition
	))

	return r
}

// return router for API endpoints
func apiRouter() http.Handler {
	r := chi.NewRouter()

	r.Post("/encrypt", handlers.Encrypt)
	r.Post("/sign", handlers.Sign)

	return r
}
