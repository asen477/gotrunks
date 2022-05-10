package middleware

import (
	"gotrunks/gin-demo/core/log"

	"github.com/gin-gonic/gin"
)

func flushLog(ctx *gin.Context) {
	log.LogSync()
}
