package http

import (
	"encoding/json"
	"net/http"
	"tinyURL/internal/shortner"
)

type Handler struct {
	svc *shortner.Service
}

func NewHandler(svc *shortner.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	code, err := h.svc.Shorten(req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]string{"code": code})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
}

func (h *Handler) ExpandURL(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[len("/"):]

	url, err := h.svc.Expand(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]string{"url": url})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
}
