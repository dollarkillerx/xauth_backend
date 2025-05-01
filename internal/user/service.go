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
	role := common.CtxGet(ctx, "role")
	fmt.Println(role)
	return &user.UserInfoResponse{
		Name:        "this is name",
		NameKana:    "kana",
		Avatar:      "",
		Email:       "",
		Phone:       "",
		Role:        "",
		AuditStatus: "",
		StudentInfo: nil,
		TeacherInfo: nil,
	}, nil
}

func (u *UserService) UpdateUserAvatar(ctx context.Context, request *user.UpdateUserAvatarRequest) (*user.UpdateUserAvatarResponse, error) {
	//TODO implement me
	panic("implement me")
}
