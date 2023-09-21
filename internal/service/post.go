package service

import (
	"context"
	"fmt"
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
	fmt.Println("postService::GetById called")
	s.store.PostStore.GetById(1)
	return nil, nil
}

func (s *postService) GetAllAfter(ctx context.Context, timestamp time.Time) ([]*dto.Post, error) {
	fmt.Println("postService::GetAllAfter called")
	s.store.PostStore.GetAllAfter(time.Now())
	return nil, nil
}
