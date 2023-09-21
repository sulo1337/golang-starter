package api

import (
	"github.com/sulo1337/cleanarch-go/internal/service"
	"log/slog"
	"net/http"
)

type UserAPI struct {
	logger      *slog.Logger
	userService service.UserService
}

func NewUserAPI(logger *slog.Logger, s service.UserService) *UserAPI {
	return &UserAPI{logger: logger, userService: s}
}

func (u *UserAPI) getByUsername(w http.ResponseWriter, r *http.Request) {
	u.logger.Debug("userRouter::getByUsername")
	u.userService.GetByUsername(r.Context(), "")
	panic("implement me")
}
