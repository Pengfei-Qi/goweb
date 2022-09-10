package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	gorm.Model
	Name string
	Age  int
}

type Result struct {
	ID   int
	Name string
	Age  int
}

/**
gorm 更新 和 删除
*/

func main() {
	db, err := gorm.Open("mysql", "root:123456@(192.168.235.233:3306)/db1?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	db.SingularTable(true)

	//自适应创建表
	db.AutoMigrate(&UserInfo{})
	//查询所有数据
	defer queryAllData(db, "user_info")

	//--------------------------------更新 begin--------------------------------

	//1. 更新单个属性，如果它有变化
	//var user UserInfo
	//db.Where("name = ?", "non_existing1").First(&user)
	//db.Debug().Model(&user).Update("name", "雨水")
	//UPDATE `user_info` SET `name` = '雨水', `updated_at` = '2022-09-08 18:02:30'  WHERE `user_info`.`deleted_at` IS NULL AND `user_info`.`id` = 8

	//2. 根据给定的条件更新单个属性
	//var user UserInfo
	//db.Debug().Model(&user).Where("name = ? ", "non_existing2").Update("name", "惊蛰")

	//3. 使用 map 更新多个属性，只会更新其中有变化的属性
	//var user UserInfo
	//db.Where("id = ?", 10).First(&user)
	//db.Debug().Model(&user).Updates(map[string]interface{}{"name": "清明", "age": 27})

	//4. 使用 struct 更新多个属性，只会更新其中有变化且为非零值的字段
	//var user UserInfo
	//db.Where("id = ?", 10).First(&user)
	//db.Model(&user).Updates(UserInfo{Name: "谷雨", Age: 28})

	//5. 警告：当使用 struct 更新时，GORM只会更新那些非零值的字段
	//var user UserInfo
	//db.Where("id = ?", 13).First(&user)
	//db.Debug().Model(&user).Updates(UserInfo{Name: "", Age: 0})

	//更新选定字段
	//6. 如果你想更新或忽略某些字段，你可以使用 Select，Omit
	//var user UserInfo
	//db.Where("id = ?", 13).First(&user)
	//db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "立夏", "age": 29})

	//6.2 Omit 跳过
	//var user UserInfo
	//db.Where("id = ?", 13).First(&user)
	//db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "小满", "age": 29})

	//无 Hooks 更新
	/**
	上面的更新操作会自动运行 model 的 BeforeUpdate, AfterUpdate 方法，更新 UpdatedAt 时间戳,
	在更新时保存其 Associations, 如果你不想调用这些方法，你可以使用 UpdateColumn， UpdateColumns
	*/

	//7.更新单个属性，类似于 `Update`
	//var user UserInfo
	//db.Where("id = ?", 12).Find(&user)
	//db.Debug().Model(&user).UpdateColumn("name", "芒种")
	// UPDATE `user_info` SET `name` = '芒种'  WHERE `user_info`.`deleted_at` IS NULL AND `user_info`.`id` = 12

	//8. 更新多个属性，类似于 `Updates`
	//var user UserInfo
	//db.Where("id = ?", 13).Find(&user)
	//db.Debug().Model(&user).UpdateColumns(&UserInfo{Name: "夏至", Age: 30})

	//批量更新
	//9. 批量更新时 Hooks 不会运行
	//db.Debug().Table("user_info").Where("id IN (?)", []int{10, 12, 13}).Updates(map[string]interface{}{"age": 88})
	//UPDATE `user_info` SET `age` = 88  WHERE (id IN (10,12,13))

	//10. 使用 struct 更新时，只会更新非零值字段，若想更新所有字段，请使用map[string]interface{}
	//db.Debug().Model(UserInfo{}).Updates(UserInfo{Age: 25})
	//UPDATE `user_info` SET `age` = 25, `updated_at` = '2022-09-08 18:40:27'  WHERE `user_info`.`deleted_at` IS NULL

	//11. 使用 `RowsAffected` 获取更新记录总数
	//affected := db.Model(UserInfo{}).Updates(UserInfo{Age: 27}).RowsAffected
	//fmt.Printf("受影响的行数为: %d \n", affected)

	//使用 SQL 表达式更新
	//12
	//var user UserInfo
	//db.Where("id = ?", 10).Find(&user)
	//db.Debug().Model(&user).Update("age", gorm.Expr("age * ? + ?", 2, 10))
	//UPDATE `user_info` SET `age` = age * 2 + 10, `updated_at` = '2022-09-08 18:47:59'  WHERE `user_info`.`deleted_at` IS NULL AND `user_info`.`id` = 10

	//13
	//var user UserInfo
	//db.Where("id = ?", 12).Find(&user)
	//db.Debug().Model(&user).Updates(map[string]interface{}{"age": gorm.Expr("age * ? + ?", 3, 10)})
	//UPDATE `user_info` SET `age` = age * 3 + 10, `updated_at` = '2022-09-08 18:51:06'  WHERE `user_info`.`deleted_at` IS NULL AND `user_info`.`id` = 12

	//14
	//var user UserInfo
	//db.Where("id = ?", 13).Find(&user)
	//db.Debug().Model(&user).Update("age", gorm.Expr("age - ?", 20))
	//UPDATE `user_info` SET `age` = age - 20, `updated_at` = '2022-09-08 18:52:55'  WHERE `user_info`.`deleted_at` IS NULL AND `user_info`.`id` = 13

	//15
	//var user []UserInfo
	//db.Debug().Model(&user).Where("id <> ?", 1).UpdateColumn("age", gorm.Expr("age + ?", 20))
	//UPDATE `user_info` SET `age` = age + 20  WHERE `user_info`.`deleted_at` IS NULL AND ((id <> 1))

	//--------------------------------更新 end--------------------------------

	//--------------------------------删除 begin--------------------------------

	/**
	警告 删除记录时，请确保主键字段有值，GORM 会通过主键去删除记录，如果主键为空，GORM 会删除该 model 的所有记录。
	*/

	//1. 删除指定数据
	//var user UserInfo
	//db.Where("id = ?", 12).Find(&user)
	//db.Debug().Delete(&user)
	//UPDATE `user_info` SET `deleted_at`='2022-09-08 19:07:55'  WHERE `user_info`.`deleted_at` IS NULL AND `user_info`.`id` = 12

	//批量删除

	//2.1 软删除-删除全部匹配的记录
	//db.Debug().Where("name like ?", "%夏%").Delete(UserInfo{})
	//UPDATE `user_info` SET `deleted_at`='2022-09-08 19:11:04'  WHERE `user_info`.`deleted_at` IS NULL AND ((name like '%夏%'))

	//2.2 软删除
	//db.Debug().Delete(UserInfo{}, "age > ?", 80)
	//UPDATE `user_info` SET `deleted_at`='2022-09-08 19:12:58'  WHERE `user_info`.`deleted_at` IS NULL AND ((age > 80))

	//3. 物理删除
	//var user UserInfo
	//db.Where("id = ?", 1).Find(&user)
	//db.Debug().Unscoped().Delete(&user)

}

func queryAllData(db *gorm.DB, tableName string) {
	var results []Result
	db.Table(tableName).Scan(&results)
	for inx, result := range results {
		fmt.Printf("查询所有 数据[遍历]  %d...%v \n", inx, result)
	}
}
