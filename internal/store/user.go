package store

import (
	"github.com/sulo1337/cleanarch-go/internal/domain"
	"gorm.io/gorm"
)

type UserStore interface {
	GetById(id uint64) (*domain.User, error)
	GetByUsername(username string) (*domain.User, error)
}

type userStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) UserStore {
	return &userStore{db: db}
}

func (s *userStore) GetById(id uint64) (*domain.User, error) {
	panic("not yet implemented")
}

func (s *userStore) GetByUsername(username string) (*domain.User, error) {
	panic("not yet implemented")
}