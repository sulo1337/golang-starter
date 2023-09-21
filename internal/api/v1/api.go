package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sulo1337/cleanarch-go/internal/service"
)

var prefix = "/api/v1"

type API struct {
	h http.Handler
	s service.Service
}

func NewAPI(s service.Service) *API {
	r := chi.NewRouter()
	
	r.Route(prefix, func(r chi.Router) {
		r.Mount("/users", userRouter())
		r.Mount("/posts", postRouter())
	})

	return &API{h: r, s: s}
}

func (a API) GetBaseRouter() http.Handler {
	return a.h
}
