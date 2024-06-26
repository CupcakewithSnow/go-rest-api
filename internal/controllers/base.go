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

	r.Post("/short", c.shortUrl)

	return r
}

func (c *BaseController) shortUrl(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var request ShortUrlRequest
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	shortUrl := request.Url + "_short"

	response := ShortUrlResponse{Url: shortUrl}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonResponse))

}
