package src

import (
	"net/http"

	"github.com/cateiru/go-http-error/httperror"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	File    string `json:"file"`
	Line    int    `json:"line"`
}

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

// ginでエラー処理を"イイカンジ"にするやつ
func (h *Handler) HandleWrapper(hand func(ctx *gin.Context) error) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		err := hand(ctx)
		if err != nil {

			if castedErr, ok := httperror.CastHTTPError(err); ok {
				ctx.AbortWithStatusJSON(castedErr.StatusCode,
					ErrorResponse{
						Status:  castedErr.StatusCode,
						Message: castedErr.Err.Error(),
						File:    castedErr.FileName,
						Line:    castedErr.Line,
					},
				)
			}
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError,
				ErrorResponse{
					Status:  http.StatusInternalServerError,
					Message: err.Error(),
					File:    "",
					Line:    0,
				})
		}
	}

}

// `/` のhandler
// とりあえず接続確認用
func (h *Handler) RootHandler(ctx *gin.Context) error {
	ctx.String(http.StatusOK, "Hello World")

	return nil
}
