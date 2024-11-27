package router

import (
	"gemini-go/internal/gemini"
	"github.com/gin-gonic/gin"
	"net/http"
)

var r *gin.Engine

func InitRouter(geminiHandler *gemini.Handler) {
	r = gin.Default()

	r.POST("/query", geminiHandler.GetResponse)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from gemini-go-serv"})
	})
}

func Start(addr string) error {
	return r.Run(addr)
}
