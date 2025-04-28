package user

import (
	"context"
	"github.com/dollarkillerx/xauth_backend/api/user"
	"github.com/dollarkillerx/xauth_backend/internal/conf"
	"github.com/dollarkillerx/xauth_backend/internal/storage"
)

type UserService struct {
	Storage *storage.Storage
	Conf    *conf.Config

	user.UnimplementedUserServiceServer
}

func (u *UserService) Register(ctx context.Context, request *user.RegisterRequest) (*user.RegisterResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) Login(ctx context.Context, request *user.LoginRequest) (*user.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) GetUserInfo(ctx context.Context, request *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	//TODO implement me
	panic("implement me")
}
