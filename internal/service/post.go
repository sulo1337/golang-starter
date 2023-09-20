package service

import (
	"context"
	"time"

	"github.com/sulo1337/cleanarch-go/internal/dto"
	"github.com/sulo1337/cleanarch-go/internal/store"
)

type PostService interface {
	GetById(ctx context.Context, id uint64) (*dto.Post, error)
	GetAllAfter(ctx context.Context, timestamp time.Time) ([]*dto.Post, error)
}

type postService struct {
	store *store.Store
}

func NewPostService(store *store.Store) PostService {
	return &postService{store: store}
}

func (s *postService) GetById(ctx context.Context, id uint64) (*dto.Post, error) {
	panic("not yet implemented")
}

func (s *postService) GetAllAfter(ctx context.Context, timestamp time.Time) ([]*dto.Post, error) {
	panic("not yet implemented")
}
