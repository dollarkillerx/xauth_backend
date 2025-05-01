package common

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"google.golang.org/grpc/peer"
	"net"
)

func CtxGet(ctx context.Context, key string) string {
	role, _ := ctx.Value(key).(string)
	return role
}

func GetClientIP(ctx context.Context) string {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return ""
	}

	if p.Addr == net.Addr(nil) {
		return ""
	}

	// 提取 IP 地址部分
	addr := p.Addr.String()
	if ip, _, err := net.SplitHostPort(addr); err == nil {
		return ip
	}
	return addr // 如果不是 host:port 格式，返回原始地址
}

const salt = "39dbf513-e548-4add-ade6-f94d49379220"

func HashPassword(password string) string {
	return Md5(password + salt)
}

func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}
