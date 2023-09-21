package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func postRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/{id}", getById)
	r.Group(func(r chi.Router) {
		r.Use(MustBeAuthenticated)
		r.Get("/", getAllAfter)
	})
	return r
}

func getById(w http.ResponseWriter, r *http.Request) {

}

func getAllAfter(w http.ResponseWriter, r *http.Request) {

}