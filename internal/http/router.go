package http

import (
	"github.com/gorilla/mux"
	"tinyURL/internal/shortner"
)

func NewRouter(svc *shortner.Service) *mux.Router {
	router := mux.NewRouter()
	handler := NewHandler(svc)

	router.HandleFunc("/shorten", handler.ShortenURL).Methods("POST")
	router.HandleFunc("/{code}", handler.ExpandURL).Methods("GET")

	return router
}
