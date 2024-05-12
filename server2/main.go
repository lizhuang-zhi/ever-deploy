package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "提供v2版本服务",
		})
	})
	r.Run("0.0.0.0:11111") // 默认监听并在 0.0.0.0:11110 上启动服务
}
