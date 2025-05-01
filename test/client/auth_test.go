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
	//resp, err := client.UserInfo(ctx, &user.UserInfoRequest{})
	//if err != nil {
	//	s, ok := status.FromError(err)
	//	if ok {
	//		fmt.Printf("错误信息: %s\n", s.Message())
	//	}
	//	log.Fatalf("调用失败: %v", err)
	//}
	//fmt.Printf("响应: %+v\n", resp)

	//client.Login(ctx, &user.LoginRequest{})

	student, err := client.RegisterStudent(ctx, &user.RegisterStudentRequest{
		Name:                     "王燁",    // 中文姓名
		NameKana:                 "オウ ヨウ", // 姓名カナ
		Avatar:                   "",
		Email:                    "leo@example.com",
		Phone:                    "+819012345678", // 日本手机号 e164 格式
		Password:                 "securePass123", // 至少 6 位
		DeviceId:                 "device-uuid-001",
		StudentNumber:            "20250001",
		Department:               "情報学部",                              // 学部（例如：信息学部）
		Major:                    "人工知能学科",                            // 学科
		EnrollmentDate:           time.Now().AddDate(-1, 0, 0).Unix(), // 1年前入学
		GraduationDate:           time.Now().AddDate(3, 0, 0).Unix(),  // 3年后毕业
		Birthday:                 time.Date(2004, 4, 12, 0, 0, 0, 0, time.UTC).Unix(),
		Gender:                   "male",
		BirthPlace:               "中国 四川省成都市",
		Address:                  "京都府京都市左京区○○町123",
		EmergencyContactName:     "王芳",
		EmergencyContactPhone:    "+8613800138000", // 中国紧急联系人
		EmergencyContactRelation: "母親",
		EmergencyContactAddress:  "中国 四川省成都市青羊区 XX路100号",
	})

	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			fmt.Printf("错误信息: %s\n", s.Message())
		}
		log.Fatalf("调用失败: %v", err)
	}
	fmt.Printf("响应: %+v\n", student)

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
