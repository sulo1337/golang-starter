package service

import (
	"context"
	"log/slog"

	"github.com/sulo1337/cleanarch-go/internal/dto"
	"github.com/sulo1337/cleanarch-go/internal/store"
)

type UserService interface {
	GetById(ctx context.Context, id uint64) (*dto.User, error)
	GetByUsername(ctx context.Context, username string) (*dto.User, error)
}

type userService struct {
	userStore store.UserStore
	logger    *slog.Logger
}

func NewUserService(logger *slog.Logger, store store.UserStore) UserService {
	return &userService{logger: logger, userStore: store}
}

func (s *userService) GetById(ctx context.Context, id uint64) (*dto.User, error) {
	s.logger.Debug("userService::GetById")
	s.userStore.GetById(id)
	return nil, nil
}

func (s *userService) GetByUsername(ctx context.Context, username string) (*dto.User, error) {
	s.logger.Debug("userService::GetByUsername")
	s.userStore.GetByUsername(username)
	return nil, nil
}
