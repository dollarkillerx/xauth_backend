package server

import (
	"github.com/dollarkillerx/xauth_backend/api/setting"
	"github.com/dollarkillerx/xauth_backend/api/user"
	"github.com/dollarkillerx/xauth_backend/internal/conf"
	"github.com/dollarkillerx/xauth_backend/internal/middleware"
	settingServer "github.com/dollarkillerx/xauth_backend/internal/setting"
	"github.com/dollarkillerx/xauth_backend/internal/storage"
	userServer "github.com/dollarkillerx/xauth_backend/internal/user"
	"google.golang.org/grpc"

	"fmt"
	"log"
	"net"
)

type Server struct {
	storage *storage.Storage
	conf    *conf.Config
}

func NewServer(storage *storage.Storage, conf *conf.Config) *Server {
	return &Server{storage: storage, conf: conf}
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.conf.ServiceConfiguration.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.GetUnaryServerInterceptor(s.conf.JWTSecretKey)),
	)

	// Register service user
	user.RegisterUserServiceServer(server, &userServer.UserService{
		Storage: s.storage,
		Conf:    s.conf,
	})

	// Register service setting
	setting.RegisterSettingServiceServer(server, &settingServer.SettingService{
		Storage: s.storage,
		Conf:    s.conf,
	})

	log.Println("gRPC server is running on port " + s.conf.ServiceConfiguration.Port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}
