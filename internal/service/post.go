package service

import (
	"context"
	"fmt"
	"github.com/sulo1337/cleanarch-go/pkg/logger"
	"time"

	"github.com/sulo1337/cleanarch-go/internal/dto"
	"github.com/sulo1337/cleanarch-go/internal/store"
)

type PostService interface {
	GetById(ctx context.Context, id uint64) (*dto.Post, error)
	GetAllAfter(ctx context.Context, timestamp time.Time) ([]*dto.Post, error)
}

type postService struct {
	postStore store.PostStore
	logger    logger.Logger
}

func NewPostService(logger logger.Logger, postStore store.PostStore) PostService {
	return &postService{logger: logger, postStore: postStore}
}

func (s *postService) GetById(ctx context.Context, id uint64) (*dto.Post, error) {
	fmt.Println("postService::GetById called")
	s.postStore.GetById(1)
	return nil, nil
}

func (s *postService) GetAllAfter(ctx context.Context, timestamp time.Time) ([]*dto.Post, error) {
	fmt.Println("postService::GetAllAfter called")
	s.postStore.GetAllAfter(time.Now())
	return nil, nil
}
