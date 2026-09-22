package main

import (
	"net/http"

	"github.com/jeromechua-12/movie-reservation/internal/helper"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	err := helper.WriteJSON(w, http.StatusOK, data, nil)
	if err != nil {
		helper.ServerErrorResponse(w, r, app.logger, err)
	}
}
