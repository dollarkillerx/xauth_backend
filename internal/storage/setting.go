package storage

import (
	"github.com/dollarkillerx/xauth_backend/pkg/models"
	"github.com/google/uuid"
)

func (s *Storage) GetSetting() (*models.Setting, error) {
	var setting models.Setting
	err := s.db.Model(&models.Setting{}).First(&setting).Error
	return &setting, err
}

func (s *Storage) UpdateSetting(setting *models.Setting) error {
	var set models.Setting
	err := s.db.Model(&models.Setting{}).First(&set).Error
	if err == nil {
		return s.db.Model(&models.Setting{}).Where("id = ?", set.ID).Updates(setting).Error
	}

	setting.BaseModel = models.BaseModel{ID: uuid.New().String()}
	return s.db.Model(&models.Setting{}).Create(setting).Error
}
