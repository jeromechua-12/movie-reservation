package main

import (
	"net/http"

	"github.com/jeromechua-12/movie-reservation/internal/auth"
)

func (app *application) routes(authHandler *auth.Handler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/healthcheck", app.healthcheckHandler)
	mux.HandleFunc("POST /v1/auth/register", authHandler.RegisterCustomer)

	return mux
}
