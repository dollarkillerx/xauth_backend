package storage

import (
	"fmt"
	"time"

	"github.com/dollarkillerx/xauth_backend/internal/user/dao"
	"github.com/dollarkillerx/xauth_backend/pkg/common"
	"github.com/dollarkillerx/xauth_backend/pkg/models"
	"github.com/google/uuid"
)

func (s *Storage) RegisterStudent(params *dao.RegisterStudentRequest) error {
	userID := uuid.New().String()

	// 创建 User 记录
	user := &models.User{
		BaseModel: models.BaseModel{
			ID: userID,
		},
		Name:               params.Name,
		NameKana:           params.NameKana,
		Avatar:             params.Avatar,
		Email:              params.Email,
		Phone:              params.Phone,
		Password:           common.HashPassword(params.Password),
		Role:               "student", // 固定为学生角色
		AuditStatus:        "pending", // 默认待审核状态
		LastLoginAt:        time.Time{},
		LastLoginDeviceID:  "",   // 默认为空 (首次登录设备 绑定后更新)
		ForbidChangeDevice: true, // 默认禁止切换设备
	}

	if err := s.db.Create(user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	// 创建 Student 记录
	student := &models.Student{
		BaseModel: models.BaseModel{
			ID: userID, // 与 user 表一一对应
		},
		StudentNumber:            params.StudentNumber,
		Department:               params.Department,
		Major:                    params.Major,
		EnrollmentDate:           &params.EnrollmentDate,
		GraduationDate:           &params.GraduationDate, // 使用指针以允许 NULL
		Birthday:                 &params.Birthday,
		Gender:                   params.Gender,
		BirthPlace:               params.BirthPlace,
		Address:                  params.Address,
		EmergencyContactName:     params.EmergencyContactName,
		EmergencyContactPhone:    params.EmergencyContactPhone,
		EmergencyContactRelation: params.EmergencyContactRelation,
		EmergencyContactAddress:  params.EmergencyContactAddress,
	}

	if err := s.db.Create(student).Error; err != nil {
		return fmt.Errorf("failed to create student: %w", err)
	}

	return nil
}

func (s *Storage) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	if err := s.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
