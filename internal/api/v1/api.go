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
	h       http.Handler
	postAPI *PostAPI
	userAPI *UserAPI
}

func NewAPI(
	logger *slog.Logger,
	postService service.PostService,
	userService service.UserService,
) *API {
	api := &API{
		postAPI: NewPostAPI(logger, postService),
		userAPI: NewUserAPI(logger, userService),
	}
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(chiMiddleware.CleanPath)
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Recoverer)

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
