package main

import (
	"flag"
	"strings"

	"github.com/dollarkillerx/xauth_backend/internal/conf"
	"github.com/dollarkillerx/xauth_backend/internal/server"
	"github.com/dollarkillerx/xauth_backend/internal/storage"
	"github.com/dollarkillerx/xauth_backend/pkg/common/client"
	"github.com/dollarkillerx/xauth_backend/pkg/common/config"
	"github.com/dollarkillerx/xauth_backend/pkg/logs"
	"github.com/rs/zerolog/log"
)

var configFilename string
var configDirs string

func init() {
	const (
		defaultConfigFilename = "config"
		configUsage           = "Name of the config file, without extension"
		defaultConfigDirs     = "./,./configs/"
		configDirUsage        = "Directories to search for config file, separated by ','"
	)
	flag.StringVar(&configFilename, "c", defaultConfigFilename, configUsage)
	flag.StringVar(&configFilename, "dev_config", defaultConfigFilename, configUsage)
	flag.StringVar(&configDirs, "cPath", defaultConfigDirs, configDirUsage)
}

func main() {
	flag.Parse()
	var appConfig conf.Config
	err := config.InitConfiguration(configFilename, strings.Split(configDirs, ","), &appConfig)
	if err != nil {
		panic(err)
	}

	// 基础依赖初始化
	// 初始化日志
	logs.InitLog(appConfig.LoggerConfiguration)

	// 初始化数据库
	postgresClient, err := client.PostgresClient(appConfig.PostgresConfiguration, nil)
	if err != nil {
		log.Error().Msg("Failed to connect to postgres")
		panic(err)
	}

	st := storage.NewStorage(postgresClient)
	log.Info().Msg("Storage initialized")

	ser := server.NewServer(st, &appConfig)
	if err := ser.Run(); err != nil {
		panic(err)
	}
}
