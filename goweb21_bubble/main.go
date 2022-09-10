package main

import (
	"goweb21_bubble/dao"
	"goweb21_bubble/models"
	"goweb21_bubble/routers"
)

func main() {
	//链接数据库信息
	err := dao.InitSql()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	dao.DB.AutoMigrate(&models.Todo{})
	//路由信息
	r := routers.SetupRouters()
	//启动端口
	r.Run(":9000")
}
