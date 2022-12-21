package src

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup, h *Handler) {

	r.GET("/", h.HandleWrapper(h.RootHandler))
}
