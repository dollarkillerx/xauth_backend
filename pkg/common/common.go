package common

import (
	"context"
	"crypto/md5"
	"encoding/hex"
)

func CtxGet(ctx context.Context, key string) string {
	role, _ := ctx.Value(key).(string)
	return role
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
