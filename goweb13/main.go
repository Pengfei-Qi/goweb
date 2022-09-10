package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//HTTP重定向
	r.GET("/index", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"status": "ok",
		//})

		c.Redirect(http.StatusMovedPermanently, "https://baidu.com")
	})
	//路由重定向
	r.GET("/home", func(c *gin.Context) {
		c.Request.URL.Path = "/company"
		r.HandleContext(c)
	})

	r.GET("/company", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"address": "company",
		})
	})

	r.Run(":9000")
}
