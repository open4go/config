package config

import (
	"context"
	"github.com/open4go/log"
	"github.com/spf13/viper"
)

const publicConfig = "./config/public.yaml"

// LoadConfig 加载配置
func LoadConfig(ctx context.Context, path string) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Log(ctx).Fatal(err)
	}

	// 使用 MergeInConfig 合并第二份配置
	viper.SetConfigFile(publicConfig)
	if err := viper.MergeInConfig(); err != nil {
		log.Log(ctx).Fatal(err)
	}
}
