package domain

import "time"

type Post struct {
	ID        uint64 `gorm:"primaryKey"`
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
