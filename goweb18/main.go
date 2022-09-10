package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID   int64
	Name string `gorm:"default:'初六'"`
	Age  int64
}

// Animal 使用`AnimalID`作为主键
//type Animal struct {
//	AnimalID int64 `gorm:"primary_key"`
//	Name     string
//	Age      int64
//}

// Animal Overriding Column Name
type Animal struct {
	ID   int64
	Name string        `gorm:"default:'galeone'"`
	Age  sql.NullInt64 `gorm:"default:55"`
}

// TableName 将 Animal 的表名设置为 `profiles`
//func (Animal) TableName() string {
//	return "wulala"
//}

func main() {
	db, err := gorm.Open("mysql", "root:123456@(192.168.235.233:3306)/db1?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	//db.SingularTable(true)

	// 使用User结构体创建名为`deleted_users`的表
	//db.Table("deleted_animal").CreateTable(&Animal{})

	// 自适应表结构
	//db.AutoMigrate(&User{})
	//db.AutoMigrate(&Animal{})

	//db.Create(&User{
	//	Name: "张三",
	//})
	//
	//var user User
	//
	//db.First(&user, 1)
	//fmt.Printf("update before user is %v \n", &user)
	//db.Model(&user).Update("updated_at", time.Now())
	//
	//db.First(&user, 1)
	//fmt.Printf("update after user is %v\n ", &user)

	// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	db.SingularTable(true)

	db.AutoMigrate(&Animal{})
	//user := User{Name: "加多宝", Age: 26}
	//fmt.Printf("主键创建值: %t \n", db.NewRecord(&user))
	//db.Create(&user)
	//fmt.Printf("主键创建值: %t \n", db.NewRecord(&user))

	//设置默认值
	var animal = Animal{Name: "赵巴拉", Age: sql.NullInt64{88, true}}
	db.Debug().Create(&animal)
}
