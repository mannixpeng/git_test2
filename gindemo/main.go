package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	e := gin.New()
	e.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
		return
	})
	e.Run("127.0.0.1:8081")
}
