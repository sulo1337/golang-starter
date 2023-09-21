package api

import (
	"fmt"
	"github.com/sulo1337/cleanarch-go/internal/service"
	"github.com/sulo1337/cleanarch-go/pkg/logger"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserAPI struct {
	logger      logger.Logger
	userService service.UserService
}

func NewUserAPI(logger logger.Logger, s service.UserService) *UserAPI {
	return &UserAPI{logger: logger, userService: s}
}

func (u *UserAPI) userRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/{username}", u.getByUsername)
	return r
}

func (u *UserAPI) getByUsername(w http.ResponseWriter, r *http.Request) {
	fmt.Println("userRouter::getByUsername called")
	u.userService.GetByUsername(r.Context(), "")
	panic("implement me")
}
