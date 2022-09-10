package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	//创建引擎
	engine := gin.Default()

	//加载静态文件
	engine.Static("/sts", "./statics")

	//自定义函数模板
	engine.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//加载引擎模板
	engine.LoadHTMLGlob("templates/**/*")
	//设置请求方式
	engine.GET("/users/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "张三",
		})
	})
	engine.GET("/posts/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "<a href='http://image.uc.cn/s/wemedia/s/upload/2022/be68d16f804d0f07b0b57a2c9dbd9fcd.png'>显示图片</a>",
		})
	})
	//使用网上模板进行展示
	engine.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})
	//运行
	err := engine.Run(":9000")
	if err != nil {
		fmt.Printf("HTTP server run failed! err: %v", err)
		return
	}
}
