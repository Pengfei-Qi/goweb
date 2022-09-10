package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	r.LoadHTMLFiles("login.html", "index.tmpl")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		//第一种取值方式
		username := c.PostForm("username")
		//第二种取值方式
		//username, _ := c.GetPostForm("username")
		password, ok := c.GetPostForm("password")
		if !ok {
			password = "***"
		}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"username": username,
			"password": password,
		})
	})

	r.Run(":9001")
}
