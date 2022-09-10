package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Userinfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("login.html")

	//返回登陆页面
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	//测试GET请求
	r.GET("/user", func(c *gin.Context) {
		var u Userinfo
		error := c.ShouldBind(&u)
		if error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": error.Error(),
			})
		} else {
			fmt.Printf("userinfo is %#v \n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	//测试post请求, 表单数据
	r.POST("/form", func(c *gin.Context) {
		var u Userinfo
		error := c.ShouldBind(&u)
		if error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": error.Error(),
			})
		} else {
			fmt.Printf("userinfo is %#v \n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	//测试post请求, json数据
	r.POST("/json", func(c *gin.Context) {
		var u Userinfo
		error := c.ShouldBind(&u)
		if error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": error.Error(),
			})
		} else {
			fmt.Printf("userinfo is %#v \n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	r.Run(":9000")
}
