package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//携带请求参数
	r.GET("/hello", func(c *gin.Context) {
		//1. 直接获取参数
		//status := c.Query("status")
		//2. 使用默认值
		//status := c.DefaultQuery("status", "这儿没有值")
		//3. 使用判断的方法
		status, ok := c.GetQuery("status")
		if !ok {
			status = "找不到值了"
		}

		c.JSON(http.StatusOK, gin.H{
			"status": status,
		})
	})

	//运行
	r.Run(":9000")
}
