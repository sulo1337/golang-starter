package api

import (
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/sulo1337/cleanarch-go/internal/api/middleware"
	v1 "github.com/sulo1337/cleanarch-go/internal/api/v1"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sulo1337/cleanarch-go/internal/service"
)

type API struct {
	h          http.Handler
	middleware middleware.Middleware
	v1         *v1.V1API
}

func NewAPI(
	logger *slog.Logger,
	postService service.PostService,
	userService service.UserService,
) *API {
	api := &API{
		middleware: middleware.NewMiddleware(logger),
	}
	r := chi.NewRouter()

	r.Use(api.middleware.RequestID)
	r.Use(api.middleware.Logger)
	r.Use(chiMiddleware.CleanPath)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Recoverer)
	api.v1 = v1.NewV1API(logger, r, api.middleware, postService, userService)
	api.h = r
	return api
}
func (a *API) GetBaseRouter() http.Handler {
	return a.h
}
