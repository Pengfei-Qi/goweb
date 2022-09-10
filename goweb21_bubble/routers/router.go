package routers

import (
	"github.com/gin-gonic/gin"
	"goweb21_bubble/controller"
)

func SetupRouters() (engine *gin.Engine) {
	r := gin.Default()
	//加载静态文件
	r.Static("/static", "static")
	//加载模板文件
	r.LoadHTMLGlob("templates/*")
	//返回首页
	r.GET("/", controller.IndexHandler)

	//调用接口
	todoGroup := r.Group("/v1")

	{
		//获取全部列表
		todoGroup.GET("/todo", controller.QueryAllTodo)
		//添加数据
		todoGroup.POST("/todo", controller.CreateTodo)
		//修改数据
		todoGroup.PUT("/todo/:id", controller.UpdateTodoInfo)
		//删除数据
		todoGroup.DELETE("/todo/:id", controller.DeleteTodo)
	}

	return r
}
