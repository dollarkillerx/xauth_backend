package setting

import (
	"context"

	"github.com/dollarkillerx/xauth_backend/api/setting"
	"github.com/dollarkillerx/xauth_backend/internal/conf"
	"github.com/dollarkillerx/xauth_backend/internal/storage"
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
