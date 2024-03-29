package store

import (
	"log/slog"
	"time"

	"github.com/sulo1337/cleanarch-go/internal/domain"
	"gorm.io/gorm"
)

type PostStore interface {
	GetById(id uint64) (*domain.Post, error)
	GetAllAfter(timestamp time.Time) ([]*domain.Post, error)
}

type postStore struct {
	logger *slog.Logger
	db     *gorm.DB
}

func NewPostStore(logger *slog.Logger, db *gorm.DB) PostStore {
	return &postStore{logger: logger, db: db}
}

func (s *postStore) GetById(id uint64) (*domain.Post, error) {
	s.logger.Debug("postStore::GetById")
	return nil, nil
}

func (s *postStore) GetAllAfter(timestamp time.Time) ([]*domain.Post, error) {
	s.logger.Debug("postStore::GetAllAfter")
	return nil, nil
}
