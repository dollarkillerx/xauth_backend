package conf

import "github.com/dollarkillerx/xauth_backend/pkg/common/config"

type Config struct {
	ServiceConfiguration  config.ServiceConfiguration
	PostgresConfiguration config.PostgresConfiguration
	LoggerConfiguration   config.LoggerConfig

	JWTSecretKey string
}
