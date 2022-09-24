package routers

import (
	"goweb30_app_framework2/logger"
	"goweb30_app_framework2/settings"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.New()

	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":     "hello app framework 简化框架结构",
			"version": settings.Conf.Version,
		})
	})

	return r
}
