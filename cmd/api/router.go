package main

import (
	"net/http"

	"github.com/jeromechua-12/movie-reservation/internal/auth"
	"github.com/jeromechua-12/movie-reservation/internal/movies"
)

func (app *application) routes(authHandler *auth.Handler, movieHandler *movies.Handler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/healthcheck", app.healthcheckHandler)

	// auth
	mux.HandleFunc("POST /v1/auth/register", authHandler.RegisterCustomer)
	mux.HandleFunc("POST /v1/auth/authenticate", authHandler.AuthenticateUser)

	// movies
	mux.HandleFunc("POST /v1/movies/create", movieHandler.CreateMovie)

	return mux
}
