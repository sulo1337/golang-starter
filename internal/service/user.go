package service

import (
	"context"
	"fmt"

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
	fmt.Println("userService::GetById called")
	s.store.UserStore.GetById(id)
	return nil, nil
}

func (s *userService) GetByUsername(ctx context.Context, username string) (*dto.User, error) {
	fmt.Println("userService::GetByUsername called")
	s.store.UserStore.GetByUsername(username)
	return nil, nil
}
