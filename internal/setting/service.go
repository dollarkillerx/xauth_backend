package setting

import (
	"context"
	"fmt"

	"github.com/dollarkillerx/xauth_backend/api/setting"
	"github.com/dollarkillerx/xauth_backend/internal/conf"
	"github.com/dollarkillerx/xauth_backend/internal/storage"
	"github.com/dollarkillerx/xauth_backend/pkg/common"
	"github.com/dollarkillerx/xauth_backend/pkg/models"
	"github.com/pkg/errors"
)

type SettingService struct {
	Storage *storage.Storage
	Conf    *conf.Config

	setting.UnimplementedSettingServiceServer
}

func (s *SettingService) GetSetting(ctx context.Context, request *setting.GetSettingRequest) (*setting.GetSettingResponse, error) {
	set, err := s.Storage.GetSetting()
	if err != nil {
		return nil, err
	}
	return &setting.GetSettingResponse{
		SchoolName:          set.SchoolName,
		PrincipalName:       set.PrincipalName,
		EmailDomain:         set.EmailDomain,
		AllowRegister:       set.AllowRegister,
		AllowChangeDevice:   set.AllowChangeDevice,
		RestrictClockInIp:   set.RestrictClockInIP,
		RestrictIpList:      set.RestrictIPList,
		RestrictEmailDomain: set.RestrictEmailDomain,
	}, nil
}

func (s *SettingService) SetSetting(ctx context.Context, request *setting.SetSettingRequest) (*setting.SetSettingResponse, error) {
	// check auth
	role := common.CtxGet(ctx, "role")
	if role != "admin" {
		return nil, fmt.Errorf("unauthorized")
	}
	// check auth end

	err := s.Storage.UpdateSetting(&models.Setting{
		SchoolName:          request.SchoolName,
		PrincipalName:       request.PrincipalName,
		EmailDomain:         request.EmailDomain,
		AllowRegister:       request.AllowRegister,
		AllowChangeDevice:   request.AllowChangeDevice,
		RestrictClockInIP:   request.RestrictClockInIp,
		RestrictIPList:      request.RestrictIpList,
		RestrictEmailDomain: request.RestrictEmailDomain,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &setting.SetSettingResponse{}, nil
}
