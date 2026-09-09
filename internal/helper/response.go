package helper

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	js, err := json.Marshal(data)
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

func WriteError(w http.ResponseWriter, status int, message string) {
	data := map[string]string{
		"error": message,
	}

	WriteJSON(w, status, data, nil)
}

func WriteServerError(w http.ResponseWriter) {
	msg := "the server encountered a problem and could not process your request"
	WriteError(w, http.StatusInternalServerError, msg)
}
