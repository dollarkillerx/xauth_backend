package middleware

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func AuthInterceptor(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("missing metadata")
	}

	tokens := md.Get("authorization") // 获取请求头中的 authorization 字段
	if len(tokens) == 0 || !validateToken(tokens[0]) {
		return nil, errors.New("invalid token")
	}

	// 认证通过，可以在 context 中加信息
	newCtx := context.WithValue(ctx, "user", "user-id-from-token")
	return newCtx, nil
}

func validateToken(token string) bool {
	return token == "your-valid-token"
}

var noAuthPath = []string{
	"/proto.UserService/Login",
	"/proto.UserService/Register",
}

func UnaryServerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	fmt.Println("调用接口:", info.FullMethod)

	// 判断是否在不需要认证的白名单中
	if !isNoAuthPath(info.FullMethod) {
		// 如果不是白名单，先认证
		newCtx, err := AuthInterceptor(ctx)
		if err != nil {
			return nil, err
		}
		ctx = newCtx // 更新 context
	}

	// 执行真正的处理逻辑
	return handler(ctx, req)
}

// 判断方法是否在白名单中
func isNoAuthPath(method string) bool {
	for _, path := range noAuthPath {
		if path == method {
			return true
		}
	}
	return false
}
