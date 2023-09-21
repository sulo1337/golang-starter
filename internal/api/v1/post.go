package api

import (
	"fmt"
	"github.com/sulo1337/cleanarch-go/internal/api/middleware"
	"github.com/sulo1337/cleanarch-go/internal/service"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type PostAPI struct {
	s *service.Service
}

func NewPostAPI(s *service.Service) *PostAPI {
	return &PostAPI{s: s}
}

func (p *PostAPI) postRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/{id}", p.getById)
	r.Group(func(r chi.Router) {
		r.Use(middleware.MustBeAuthenticated)
		r.Get("/", p.getAllAfter)
	})
	return r
}

func (p *PostAPI) getById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("postRouter::getById called")
	p.s.PostService.GetById(r.Context(), 1)
}

func (p *PostAPI) getAllAfter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("postRouter::getAllAfter called")
	p.s.PostService.GetAllAfter(r.Context(), time.Now())
}
