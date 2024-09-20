package config

import (
	"context"
	"github.com/open4go/log"
	"github.com/spf13/viper"
)

// LoadConfig 加载配置
func LoadConfig(ctx context.Context, path string) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Log(ctx).Fatal(err)
	}
	log.Log(ctx).WithField("path", path).
		Info("loading config successful")
}
