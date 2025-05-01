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

}
