package main

import (
	"GoServe/internal/controllers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	chi := chi.NewRouter()

	chi.Mount("/", controllers.NewBaseController().Route())

	log.Fatal(http.ListenAndServe(":8080", chi))
}
