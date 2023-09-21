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

func NewAPI(
	postService service.PostService,
	userService service.UserService,
) *API {
	api := &API{
		postAPI: NewPostAPI(postService),
		userAPI: NewUserAPI(userService),
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
