package storage

import "github.com/dollarkillerx/xauth_backend/pkg/models"

func (s *Storage) CreateTeacher(params *models.Teachers) error {
	return s.db.Create(params).Error
}

func (s *Storage) GetTeacherByID(id string) (*models.Teachers, error) {
	teacher := &models.Teachers{}
	if err := s.db.Where("id = ?", id).First(teacher).Error; err != nil {
		return nil, err
	}
	return teacher, nil
}
