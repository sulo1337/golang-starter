package api

import (
	"fmt"
	"github.com/sulo1337/cleanarch-go/internal/service"
	"log/slog"
	"net/http"
	"time"
)

type PostAPI struct {
	logger      *slog.Logger
	postService service.PostService
}

func NewPostAPI(logger *slog.Logger, s service.PostService) *PostAPI {
	return &PostAPI{logger: logger, postService: s}
}

func (p *PostAPI) getById(w http.ResponseWriter, r *http.Request) {
	p.logger.Info("ctx: %v", r.Context())
	fmt.Println("postRouter::getById called")
	p.postService.GetById(r.Context(), 1)
}

func (p *PostAPI) getAllAfter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("postRouter::getAllAfter called")
	p.postService.GetAllAfter(r.Context(), time.Now())
}
