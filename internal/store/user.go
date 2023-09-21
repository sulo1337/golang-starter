package store

import (
	"fmt"
	"github.com/sulo1337/cleanarch-go/internal/domain"
	"gorm.io/gorm"
	"log/slog"
)

type UserStore interface {
	GetById(id uint64) (*domain.User, error)
	GetByUsername(username string) (*domain.User, error)
}

type userStore struct {
	logger *slog.Logger
	db     *gorm.DB
}

func NewUserStore(logger *slog.Logger, db *gorm.DB) UserStore {
	return &userStore{logger: logger, db: db}
}

func (s *userStore) GetById(id uint64) (*domain.User, error) {
	fmt.Println("userStore::GetById called")
	return nil, nil
}

func (s *userStore) GetByUsername(username string) (*domain.User, error) {
	fmt.Println("userStore::GetByUsername called")
	return nil, nil
}
