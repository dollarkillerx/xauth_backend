package test

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
	"time"

	"github.com/dollarkillerx/xauth_backend/api/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestAuth1(t *testing.T) {
	// 建立连接
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// 创建客户端连接
	conn, err := grpc.NewClient("dns:///127.0.0.1:8181", opts...)
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()

	// 初始化客户端（假设你的服务叫 UserServiceClient）
	client := user.NewUserServiceClient(conn)

	// 构造请求带上 metadata（带 Token）
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "your-valid-token1")

	// 设置超时时间
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	// 调用方法，比如调用 UserInfo
	//resp, err := client.UserInfo(ctx, &user.UserInfoRequest{})
	//if err != nil {
	//	s, ok := status.FromError(err)
	//	if ok {
	//		fmt.Printf("错误信息: %s\n", s.Message())
	//	}
	//	log.Fatalf("调用失败: %v", err)
	//}

	client.Login(ctx, &user.LoginRequest{})

	//fmt.Printf("响应: %+v\n", resp)
}
