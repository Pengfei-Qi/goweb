package models

import "goweb21_bubble/dao"

// Todo 创建模型
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func GetAllTodo(todoList *[]Todo) (err error) {
	err = dao.DB.Find(&todoList).Error
	return
}

func CreateTodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetTodoById(todoId string, todo *Todo) (err error) {
	err = dao.DB.Where("id = ?", todoId).First(&todo).Error
	return
}

func UpdateTodo(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}

func DeleteTodo(todoId string) (err error) {
	err = dao.DB.Where("id = ?", todoId).Delete(&Todo{}).Error
	return
}
