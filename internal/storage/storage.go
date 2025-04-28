package storage

import (
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB // pgsql
}

func NewStorage(db *gorm.DB) *Storage {

	return &Storage{db: db}
}

func (s *Storage) DB() *gorm.DB {
	return s.db
}
