package storage

import (
	"fmt"
	"time"

	"github.com/dollarkillerx/xauth_backend/pkg/models"
	"github.com/google/uuid"
)

func (s *Storage) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	if err := s.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Storage) UpdateUserLoginLogs(user *models.User, deviceID string, ip string) error {
	err := s.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"last_login_at":        time.Now(),
		"last_login_device_id": deviceID,
	}).Error

	if err != nil {
		return fmt.Errorf("failed to update user login logs: %w", err)
	}

	s.db.Model(&models.LoginLog{}).Create(&models.LoginLog{
		BaseModel: models.BaseModel{
			ID: uuid.New().String(),
		},
		UserID:    user.ID,
		DeviceID:  deviceID,
		IP:        ip,
		UserAgent: "",
		LoginAt:   time.Now(),
	})

	return err
}

func (s *Storage) GetUserByID(id string) (*models.User, error) {
	user := &models.User{}
	if err := s.db.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Storage) UpdateUserAvatar(userId string, avatar string) error {
	return s.db.Model(&models.User{}).Where("id = ?", userId).Updates(map[string]interface{}{
		"avatar": avatar,
	}).Error
}
