package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BaseHandler func(w http.ResponseWriter, r *http.Request) error

func (fn BaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	if err == nil {
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(&err)
}

// RespondOK writes the given data to an HTTP response with a status of 200.
func RespondOK(w http.ResponseWriter, data interface{}) error {
	if err := Respond(w, data, http.StatusOK); err != nil {
		return fmt.Errorf("respond: %w", err)
	}

	return nil
}

func RespondCreated(w http.ResponseWriter, data interface{}) error {
	if err := Respond(w, data, http.StatusCreated); err != nil {
		return fmt.Errorf("respond: %w", err)
	}

	return nil
}

func RespondBadRequest(w http.ResponseWriter, data interface{}) error {
	if err := Respond(w, data, http.StatusBadRequest); err != nil {
		return fmt.Errorf("respond: %w", err)
	}
	return nil
}

func Respond500(w http.ResponseWriter, data interface{}) error {
	if err := Respond(w, data, http.StatusInternalServerError); err != nil {
		return fmt.Errorf("respond: %w", err)
	}
	return nil
}

func RespondNotFound(w http.ResponseWriter, data interface{}) error {
	if err := Respond(w, data, http.StatusNotFound); err != nil {
		return fmt.Errorf("respond: %w", err)
	}
	return nil
}

// Respond writes the given data to an HTTP response with a status code.
func Respond(w http.ResponseWriter, data interface{}, statusCode int) error {
	w.WriteHeader(statusCode)

	// When there is no data to return to the client or the desired status code
	// is NoContent, we want to return and not write data back to the response header.
	if data == nil || statusCode == http.StatusNoContent {
		return nil
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("encode: %w", err)
	}
	return nil
}
