package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/sulo1337/cleanarch-go/internal/api/middleware"
	"github.com/sulo1337/cleanarch-go/internal/service"
	"log/slog"
)

var prefix = "/api/v1"

type V1API struct {
	middleware middleware.Middleware
	postAPI    *PostAPI
	userAPI    *UserAPI
}

func NewV1API(
	logger *slog.Logger,
	r *chi.Mux,
	m middleware.Middleware,
	postService service.PostService,
	userService service.UserService,
) *V1API {
	v1API := &V1API{
		middleware: m,
		postAPI:    NewPostAPI(logger, postService),
		userAPI:    NewUserAPI(logger, userService),
	}
	setupRoutes(r, v1API)
	return v1API
}

func setupRoutes(r *chi.Mux, api *V1API) chi.Router {
	return r.Route(prefix, func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/{username}", api.userAPI.getByUsername)
		})
		r.Route("/posts", func(r chi.Router) {
			r.With(api.middleware.Authenticated).Get("/{id}", api.postAPI.getById)
			r.Get("/", api.postAPI.getAllAfter)
		})
	})
}
