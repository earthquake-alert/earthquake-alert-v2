package src

import (
	"context"

	"github.com/earthquake-alert/erarthquake-alert-v2/src/logging"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Init(mode string) {
	logging.InitLogging(mode)
	err := InitConfig(mode)
	if err != nil {
		logging.Sugar.Fatal(err)
	}
}

func Server() {
	ctx := context.Background()
	r := gin.New()
	ServerMiddleWare(r)
	g := BasicAuth(r)

	db, err := NewConnectMySQL(ctx)
	if err != nil {
		logging.Sugar.Fatal(err)
	}
	defer db.Close()
	h := NewHandler(db)

	Routes(g, h)

	if err := r.Run(); err != nil {
		logging.Sugar.Fatal(err)
	}
}

// Ginのミドルウェア
func ServerMiddleWare(r *gin.Engine) {
	// ログをzapで出す
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
}

func BasicAuth(r *gin.Engine) *gin.RouterGroup {
	// Basic AuthがConfigで設定されていない場合は使用しない
	if C.AuthenticationUser == "" || C.AuthenticationPw == "" {
		return r.Group("/")
	}

	logging.Sugar.Info("Used BasicAuth")

	return r.Group("/", gin.BasicAuth(gin.Accounts{
		C.AuthenticationUser: C.AuthenticationPw,
	}))
}
