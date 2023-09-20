package store

import (
	"time"

	"github.com/sulo1337/cleanarch-go/internal/domain"
	"gorm.io/gorm"
)

type PostStore interface {
	GetById(id uint64) (*domain.Post, error)
	GetAllAfter(timestamp time.Time) ([]*domain.Post, error)
}

type postStore struct {
	db *gorm.DB
}

func NewPostStore(db *gorm.DB) PostStore {
	return &postStore{db: db}
}

func (s *postStore) GetById(id uint64) (*domain.Post, error) {
	panic("not yet implemented")
}

func (s *postStore) GetAllAfter(timestamp time.Time) ([]*domain.Post, error) {
	panic("not yet implemented")
}
