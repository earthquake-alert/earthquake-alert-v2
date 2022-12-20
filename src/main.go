package src

import (
	"context"
	"time"

	"github.com/earthquake-alert/erarthquake-alert-v2/src/logging"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Mode = ""

func Init(mode string) {
	Mode = mode
	logging.InitLogging(mode)
	err := InitConfig(mode)
	if err != nil {
		logging.Sugar.Fatal(err)
	}

	_, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		logging.Sugar.Fatal(err)
	}
}

func Server() {
	ctx := context.Background()
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

	db, err := NewConnectMySQL(ctx)
	if err != nil {
		logging.Sugar.Fatal(err)
	}
	h := NewHandler(db)
	Routes(r, h)

	if err := r.Run(); err != nil {
		logging.Sugar.Fatal(err)
	}
}
