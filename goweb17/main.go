package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@(192.168.235.233:3306)/db1?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	//创建表
	db.AutoMigrate(&UserInfo{})
	// 创建
	//db.Create(&UserInfo{
	//	Name:   "李思思",
	//	Gender: "女",
	//	Hobby:  "乒乓球",
	//})

	//读取数据库
	var userInfo UserInfo
	db.First(&userInfo)
	fmt.Printf("user %#v \n", userInfo)

	//更新
	//db.Model(&userInfo).Update("Name", "老麻子")

	//删除
	//db.Delete(&userInfo)
}
