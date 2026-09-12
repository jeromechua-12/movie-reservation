package helper

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func logError(r *http.Request, logger *slog.Logger, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	logger.Error(err.Error(), "method", method, "uri", uri)
}

func WriteJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, values := range headers {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func WriteError(w http.ResponseWriter, r *http.Request, logger *slog.Logger, status int, message any) {
	data := map[string]any{
		"error": message,
	}

	err := WriteJSON(w, status, data, nil)
	if err != nil {
		logError(r, logger, err)
	}
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	msg := "the server encountered a problem and could not process your request"
	WriteError(w, r, logger, http.StatusInternalServerError, msg)
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, logger *slog.Logger, err error) {
	WriteError(w, r, logger, http.StatusBadRequest, err.Error())
}

func FailedValidationResponse(w http.ResponseWriter, r *http.Request, logger *slog.Logger, errors map[string]string) {
	WriteError(w, r, logger, http.StatusUnprocessableEntity, errors)
}
