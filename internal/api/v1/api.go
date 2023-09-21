package api

import (
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/sulo1337/cleanarch-go/internal/api/middleware"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sulo1337/cleanarch-go/internal/service"
)

var prefix = "/api/v1"

type API struct {
	h          http.Handler
	middleware middleware.Middleware
	postAPI    *PostAPI
	userAPI    *UserAPI
}

func NewAPI(
	logger *slog.Logger,
	postService service.PostService,
	userService service.UserService,
) *API {
	api := &API{
		middleware: middleware.NewMiddleware(logger),
		postAPI:    NewPostAPI(logger, postService),
		userAPI:    NewUserAPI(logger, userService),
	}
	r := chi.NewRouter()

	r.Use(api.middleware.RequestID)
	r.Use(chiMiddleware.CleanPath)
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Recoverer)

	setupRoutes(r, api)

	api.h = r
	return api
}

func setupRoutes(r *chi.Mux, api *API) chi.Router {
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

func (a *API) GetBaseRouter() http.Handler {
	return a.h
}
