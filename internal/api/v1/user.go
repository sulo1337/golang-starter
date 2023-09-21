package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func userRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/{username}", getByUsername)
	return r
}

func getByUsername(w http.ResponseWriter, r *http.Request) {

}