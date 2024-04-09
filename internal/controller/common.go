package controller

import "github.com/go-chi/chi"

const dateFormat = "2006-01-02"

// HTTP controller for HTTP protocol.
type HTTP interface {
	SetRoutes(r chi.Router)
}

// basicMessage is the representation of a basic http JSON response.
type basicMessage struct {
	Message string `json:"message"`
}
