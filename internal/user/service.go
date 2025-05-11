package user

import (
	"context"
	"fmt"
	"time"

	"github.com/dollarkillerx/xauth_backend/api/user"
	"github.com/dollarkillerx/xauth_backend/internal/conf"
	"github.com/dollarkillerx/xauth_backend/internal/storage"
	"github.com/dollarkillerx/xauth_backend/internal/user/dao"
	"github.com/dollarkillerx/xauth_backend/pkg/common"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	Storage *storage.Storage
	Conf    *conf.Config

	user.UnimplementedUserServiceServer
}

func (u *UserService) RegisterStudent(ctx context.Context, request *user.RegisterStudentRequest) (*user.RegisterStudentResponse, error) {
	// check auth
	role := common.CtxGet(ctx, "role")
	if role != "admin" {
		return nil, fmt.Errorf("unauthorized")
	}
	// check auth end
	params := dao.NewRegisterStudentRequestFromGRPC(request)

	err := validator.New().Struct(params)
	if err != nil {
		return nil, err
	}

	err = u.Storage.RegisterStudent(params)

	return &user.RegisterStudentResponse{}, nil
}

func (u *UserService) RegisterTeacher(ctx context.Context, request *user.RegisterTeacherRequest) (*user.RegisterTeacherResponse, error) {
	// check auth
	role := common.CtxGet(ctx, "role")
	if role != "admin" {
		return nil, fmt.Errorf("unauthorized")
	}
	// check auth end

	//TODO implement me
	panic("implement me")
}

func (u *UserService) Login(ctx context.Context, request *user.LoginRequest) (*user.LoginResponse, error) {
	params := dao.NewLoginRequestFromGRPC(request)

	err := validator.New().Struct(params)
	if err != nil {
		return nil, err
	}

	userdata, err := u.Storage.GetUserByEmail(params.Email)
	if err != nil {
		return nil, err
	}

	if userdata.Password != common.HashPassword(params.Password) {
		return nil, fmt.Errorf("invalid password")
	}

	if userdata.AuditStatus != "disabled" {
		return nil, fmt.Errorf("user is disabled")
	}

	jwt, err := generateJWT(userdata.ID, userdata.Role, u.Conf.JWTSecretKey, time.Hour*24*180)
	if err != nil {
		return nil, err
	}

	// login 日志更新
	err = u.Storage.UpdateUserLoginLogs(userdata, request.DeviceId, common.GetClientIP(ctx))

	return &user.LoginResponse{
		Jwt: jwt,
	}, nil
}

func generateJWT(userID string, role string, secretKey string, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(duration).Unix(), // 过期时间
		"iat":     time.Now().Unix(),               // 签发时间
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}

func (u *UserService) UserInfo(ctx context.Context, request *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	userId := common.CtxGet(ctx, "user_id")
	userInfo, err := u.Storage.GetUserByID(userId)
	if err != nil {
		return nil, err
	}

	var studentInfo *user.StudentInfo
	var teacherInfo *user.TeacherInfo

	if userInfo.Role == "student" {
		student, err := u.Storage.GetStudentByID(userId)
		if err != nil {
			return nil, err
		}
		studentInfo = &user.StudentInfo{
			StudentNumber:            student.StudentNumber,
			Department:               student.Department,
			Major:                    student.Major,
			EnrollmentDate:           student.EnrollmentDate.Unix(),
			GraduationDate:           student.GraduationDate.Unix(),
			Birthday:                 student.Birthday.Unix(),
			Gender:                   student.Gender,
			BirthPlace:               student.BirthPlace,
			Address:                  student.Address,
			EmergencyContactName:     student.EmergencyContactName,
			EmergencyContactPhone:    student.EmergencyContactPhone,
			EmergencyContactRelation: student.EmergencyContactRelation,
			EmergencyContactAddress:  student.EmergencyContactAddress,
		}
	}

	if userInfo.Role == "teacher" {
		teacher, err := u.Storage.GetTeacherByID(userId)
		if err != nil {
			return nil, err
		}

		teacherInfo = &user.TeacherInfo{
			Introduction:  teacher.Introduction,
			TeacherNumber: teacher.TeacherNumber,
		}
	}

	return &user.UserInfoResponse{
		Name:        userInfo.Name,
		NameKana:    userInfo.NameKana,
		Avatar:      userInfo.Avatar,
		Email:       userInfo.Email,
		Phone:       userInfo.Phone,
		Role:        userInfo.Role,
		AuditStatus: userInfo.AuditStatus,
		StudentInfo: studentInfo,
		TeacherInfo: teacherInfo,
	}, nil
}

func (u *UserService) UpdateUserAvatar(ctx context.Context, request *user.UpdateUserAvatarRequest) (*user.UpdateUserAvatarResponse, error) {
	err := u.Storage.UpdateUserAvatar(common.CtxGet(ctx, "user_id"), request.Avatar)
	return &user.UpdateUserAvatarResponse{}, err
}
