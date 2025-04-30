package middleware

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// 不需要认证的路径
var noAuthPath = []string{
	"/user.UserService/Login",
}

func GetUnaryServerInterceptor(secretKey string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("call method:", info.FullMethod)

		if !isNoAuthPath(info.FullMethod) {
			// 从 ctx 提取 token 并校验
			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
			}

			tokens := md.Get("authorization")
			if len(tokens) == 0 {
				return nil, status.Errorf(codes.Unauthenticated, "missing token")
			}

			claims, err := validateJWT(tokens[0], secretKey)
			if err != nil {
				return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
			}

			// 从 claims 提取 user_id
			userID, ok := claims["user_id"].(string)
			if !ok {
				return nil, status.Errorf(codes.Unauthenticated, "invalid token payload: user_id missing")
			}

			// 注入到上下文
			ctx = context.WithValue(ctx, "user_id", userID)
		}

		return handler(ctx, req)
	}
}

func validateJWT(tokenString string, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
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
