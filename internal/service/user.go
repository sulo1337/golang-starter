package service

import (
	"context"
	"fmt"
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
	fmt.Println("userService::GetById called")
	s.userStore.GetById(id)
	return nil, nil
}

func (s *userService) GetByUsername(ctx context.Context, username string) (*dto.User, error) {
	fmt.Println("userService::GetByUsername called")
	s.userStore.GetByUsername(username)
	return nil, nil
}
