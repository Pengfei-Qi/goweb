package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//map 形式的请求
	r.GET("/json1", func(c *gin.Context) {
		data := gin.H{"name": "张三", "age": 20, "message": "hello world!"}
		c.JSON(http.StatusOK, data)
	})
	//结构体的请求方式
	type msg struct {
		Name    string
		Age     int
		Message string `json:"message"`
	}
	r.GET("/json2", func(c *gin.Context) {
		data := msg{
			"李四",
			34,
			"hello goland",
		}
		c.JSON(http.StatusOK, data)
	})
	//创建端口, 运行服务
	r.Run(":9000")
}
