package store

import (
	"fmt"
	"github.com/sulo1337/cleanarch-go/internal/domain"
	"github.com/sulo1337/cleanarch-go/pkg/logger"
	"gorm.io/gorm"
)

type UserStore interface {
	GetById(id uint64) (*domain.User, error)
	GetByUsername(username string) (*domain.User, error)
}

type userStore struct {
	logger logger.Logger
	db     *gorm.DB
}

func NewUserStore(logger logger.Logger, db *gorm.DB) UserStore {
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
