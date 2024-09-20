package config

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/open4go/log"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"github.com/spf13/viper"
)

type RouteLoader interface {
	LoadRoute(r *gin.Engine, cors, prefix string, jwtKey []byte, mode string)
}

func SetupGinMetrics(r *gin.Engine, prefix string) {
	m := ginmetrics.GetMonitor()
	m.SetMetricPath(prefix + "/metrics")
	m.SetSlowTime(10)
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	m.Use(r)
}

func LoadRoutesAndStartServer(ctx context.Context, r *gin.Engine, jwtKey []byte, loader RouteLoader) {
	cors := viper.GetString("server.cors")
	prefix := viper.GetString("server.prefix")
	mode := viper.GetString("gin.mode")

	loader.LoadRoute(r, cors, prefix, jwtKey, mode)

	port := viper.GetString("server.port")
	log.Log(ctx).WithField("port", port).
		WithField("prefix", prefix).
		Info("Server is running successfully")

	if err := r.Run(":" + port); err != nil {
		log.Log(ctx).Fatal(err)
	}
}
