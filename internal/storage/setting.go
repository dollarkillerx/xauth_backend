package storage

import "github.com/dollarkillerx/xauth_backend/pkg/models"

func (s *Storage) GetSetting() (*models.Setting, error) {
	var setting models.Setting
	err := s.db.Model(&models.Setting{}).First(&setting).Error
	return &setting, err
}
