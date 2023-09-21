package service

import (
	"context"
	"log/slog"
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
	logger    slog.Logger
}

func NewPostService(logger *slog.Logger, postStore store.PostStore) PostService {
	return &postService{logger: *logger, postStore: postStore}
}

func (s *postService) GetById(ctx context.Context, id uint64) (*dto.Post, error) {
	s.logger.Debug("postService::GetById")
	s.postStore.GetById(1)
	return nil, nil
}

func (s *postService) GetAllAfter(ctx context.Context, timestamp time.Time) ([]*dto.Post, error) {
	s.logger.Debug("postService::GetAllAfter")
	s.postStore.GetAllAfter(time.Now())
	return nil, nil
}
