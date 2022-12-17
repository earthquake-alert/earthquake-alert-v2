package src

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, h *Handler) {

	r.GET("/", h.HandleWrapper(h.RootHandler))
}
