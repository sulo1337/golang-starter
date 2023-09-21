package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sulo1337/cleanarch-go/internal/service"
)

var prefix = "/api/v1"

type API struct {
	h       http.Handler
	postAPI *PostAPI
	userAPI *UserAPI
}

func NewAPI(s *service.Service) *API {
	api := &API{
		postAPI: NewPostAPI(s),
		userAPI: NewUserAPI(s),
	}
	r := chi.NewRouter()

	r.Route(prefix, func(r chi.Router) {
		r.Mount("/users", api.userAPI.userRouter())
		r.Mount("/posts", api.postAPI.postRouter())
	})

	api.h = r
	return api
}

func (a *API) GetBaseRouter() http.Handler {
	return a.h
}
