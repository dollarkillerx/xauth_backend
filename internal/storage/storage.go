package storage

import (
	"github.com/dollarkillerx/xauth_backend/pkg/models"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB // pgsql
}

func NewStorage(db *gorm.DB) *Storage {
	db.AutoMigrate(
		&models.User{},
		&models.Teachers{},
		&models.Student{},
		&models.Notification{},
		&models.Setting{},
		&models.SchoolLocation{},
		&models.CourseEnrollment{},
		&models.Attendance{},
		&models.Courses{},
	)
	return &Storage{db: db}
}

func (s *Storage) DB() *gorm.DB {
	return s.db
}
