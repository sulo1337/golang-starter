package service

import (
	"context"

	"github.com/sulo1337/cleanarch-go/internal/dto"
	"github.com/sulo1337/cleanarch-go/internal/store"
)

type UserService interface {
	GetById(ctx context.Context, id uint64) (*dto.User, error)
	GetByUsername(ctx context.Context, username string) (*dto.User, error)
}

type userService struct {
	store *store.Store
}

func NewUserService(store *store.Store) UserService {
	return &userService{store: store}
}

func (s *userService) GetById(ctx context.Context, id uint64) (*dto.User, error) {
	panic("not yet implemented")
}

func (s *userService) GetByUsername(ctx context.Context, username string) (*dto.User, error) {
	panic("not yet implemented")
}
