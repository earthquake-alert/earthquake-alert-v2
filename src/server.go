package src

import (
	"github.com/earthquake-alert/erarthquake-alert-v2/src/logging"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Init(mode string) {
	logging.InitLogging(mode)
}

func Server() error {
	r := gin.New()

	r.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		if params.StatusCode == 200 {
			logging.Logger.Info("request",
				zap.String("method", params.Method),
				zap.String("path", params.Path),
				zap.Int("status", params.StatusCode),
				zap.String("host", params.Request.Host),
				zap.String("response_time", params.Latency.String()),
				zap.String("error_message", params.ErrorMessage),
			)
		} else {
			logging.Logger.Error("request",
				zap.String("method", params.Method),
				zap.String("path", params.Path),
				zap.Int("status", params.StatusCode),
				zap.String("host", params.Request.Host),
				zap.String("response_time", params.Latency.String()),
				zap.String("error_message", params.ErrorMessage),
			)
		}

		return ""
	}))
	r.Use(gin.Recovery())

	h := NewHandler()
	Routes(r, h)

	return r.Run()
}