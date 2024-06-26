package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type BaseController struct{}

type ShortUrlRequest struct {
	Url string `json:"url"`
}

type ShortUrlResponse struct {
	Url string `json:"url"`
}

func NewBaseController() *BaseController {
	return &BaseController{}
}

func (c *BaseController) Route() *chi.Mux {
	r := chi.NewRouter()

	r.HandleFunc("/{url}", c.shortUrl)

	return r
}

func (c *BaseController) shortUrl(w http.ResponseWriter, r *http.Request) {
	url := chi.URLParam(r, "url")

	response := ShortUrlResponse{Url: url + "_short"}

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonResponse))

}
