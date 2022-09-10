package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./404.html", "./index.html")
	//常规方式
	r.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	})
	r.POST("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"method": "POST"})
	})
	r.PUT("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"method": "PUT"})
	})
	//统配方式
	r.Any("/shop", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{"method": "shop/GET"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"method": "shop/POST"})
		case http.MethodPut:
			c.JSON(http.StatusOK, gin.H{"method": "shop/PUT"})
		default:
			c.JSON(http.StatusOK, gin.H{"method": fmt.Sprintf("未匹配到准确方式,为: %s", c.Request.Method)})
		}

	})

	//未匹配到路由
	r.NoRoute(func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{"msg": "这儿没有你想要的数据"})
		c.HTML(http.StatusOK, "404.html", nil)
	})
	//路由群组
	marketGroup := r.Group("/market")
	{
		marketGroup.GET("/m1", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": http.MethodGet,
				"url":    c.Request.URL.Path,
			})
		})
		marketGroup.GET("/m2", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": http.MethodGet,
				"url":    c.Request.URL.Path,
			})
		})
		marketGroup.POST("/m3", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": http.MethodPost,
				"url":    c.Request.URL.Path,
			})
		})
		mcnGroup := marketGroup.Group("/mcn")
		{
			mcnGroup.Any("/c1", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"method": c.Request.Method,
					"url":    c.Request.URL.Path,
				})
			})
		}
	}

	r.Run(":9000")
}
