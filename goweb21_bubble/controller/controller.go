package controller

import (
	"github.com/gin-gonic/gin"
	"goweb21_bubble/models"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func QueryAllTodo(c *gin.Context) {
	var todoList []models.Todo
	if err := models.GetAllTodo(&todoList); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	//绑定数据
	c.Bind(&todo)
	//将数据入库
	err2 := models.CreateTodo(&todo)
	if err2 != nil {
		c.JSON(http.StatusOK, gin.H{"error": err2.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "success"})
	}
}

func UpdateTodoInfo(c *gin.Context) {
	//获取id
	todoId, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"err": "id 值为空"})
		return
	}
	//更新数据库
	var todo models.Todo
	//判断数据库中是否存在
	if err2 := models.GetTodoById(todoId, &todo); err2 != nil {
		c.JSON(http.StatusOK, gin.H{"err": err2})
		return
	}

	c.Bind(&todo)
	if err3 := models.UpdateTodo(&todo); err3 != nil {
		c.JSON(http.StatusOK, gin.H{"err": err3.Error()})
		return
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(c *gin.Context) {
	//获取id
	todoId := c.Param("id")
	var todo models.Todo
	//判断数据库中是否存在
	if err2 := models.GetTodoById(todoId, &todo); err2 != nil {
		c.JSON(http.StatusOK, gin.H{"err": "找不到该数据"})
		return
	}
	//删除数据
	if err := models.DeleteTodo(todoId); err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "删除成功"})
	}
}
