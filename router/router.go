package router

import (
	"gemini-go/internal/gemini"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(geminiHandler *gemini.Handler) {
	r = gin.Default()

	r.POST("/query", geminiHandler.GetResponse)
}

func Start(addr string) error {
	return r.Run(addr)
}
