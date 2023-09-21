package store

import (
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
	s.logger.Debug("userStore::GetById")
	return nil, nil
}

func (s *userStore) GetByUsername(username string) (*domain.User, error) {
	s.logger.Debug("userStore::GetByUsername")
	return nil, nil
}
