package config

import (
	"context"
	"encoding/base64"
	"github.com/open4go/log"
	"github.com/spf13/viper"
)

const (
	// jwtKeyPath 配置
	jwtKeyPath = "jwt.key"
	// jwtEnabledPath 配置
	jwtEnabledPath = "jwt.enabled"
)

// DecodeJWTKey 解密jwt
func DecodeJWTKey(ctx context.Context, path string) []byte {
	jwtKey := viper.GetString(jwtKeyPath)
	enableJWT := viper.GetBool(jwtEnabledPath)
	if enableJWT && jwtKey != "" {
		sEnc := base64.StdEncoding.EncodeToString([]byte(jwtKey))
		sDec, err := base64.StdEncoding.DecodeString(sEnc)
		if err != nil {
			log.Log(ctx).Fatal(err)
		}
		return sDec
	}
	return nil
}
