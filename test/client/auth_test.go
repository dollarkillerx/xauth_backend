package test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/dollarkillerx/xauth_backend/api/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDYxNjEwNDksImlhdCI6MTc0NjA3NDY0OSwicm9sZSI6ImFkbWluIiwidXNlcl9pZCI6ImFkYXBhd2FuZ0BnbWFpbC5jb20ifQ.h9dJqqEk_I6nHH06MhAErZrDC9jE08MeeYCIRFa8qPo")

	// 设置超时时间
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	// 调用方法，比如调用 UserInfo
	resp, err := client.UserInfo(ctx, &user.UserInfoRequest{})
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			fmt.Printf("错误信息: %s\n", s.Message())
		}
		log.Fatalf("调用失败: %v", err)
	}
	fmt.Printf("响应: %+v\n", resp)

	//client.Login(ctx, &user.LoginRequest{})

}

func TestAuth2(t *testing.T) {
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
	login, err := client.Login(ctx, &user.LoginRequest{
		Email:    "adapawang@gmail.com",
		Password: "",
		DeviceId: "",
	})
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			fmt.Printf("错误信息: %s\n", s.Message())
		}
		log.Fatalf("调用失败: %v", err)
	}

	fmt.Printf("响应: %+v\n", login)

}
