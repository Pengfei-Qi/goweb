package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
gorm查询
*/

type UserInfo struct {
	gorm.Model
	Name string
	Age  int
}

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

	//user := UserInfo{Name: "孙权", Age: 22}
	//db.Create(&user)
	//
	//user2 := UserInfo{Name: "吴越", Age: 23}
	//db.Create(&user2)

	//user3 := UserInfo{}
	// 根据主键查询第一条记录
	//db.First(&user3)
	//fmt.Printf("查询第一条记录........%#v \n", user3)

	//随机查询一条记录
	//db.Take(&user3)
	//fmt.Printf("随机查询一条记录........%#v \n", user3)

	//根据主键查询最后一条记录
	//db.Last(&user3)
	//fmt.Printf("根据主键查询最后一条记录........%#v \n", user3)

	// 查询所有的记录
	//var user4 []UserInfo
	//db.Find(&user4)
	//for _, u := range user4 {
	//	fmt.Printf("查询所有的记录........%v \n", u)
	//}

	//Where 条件
	//var user UserInfo
	//db.Where("name = ?", "赵六").First(&user)
	//fmt.Printf("根据条件查询,获取第一条匹配的记录...%v \n", user)

	//获取所有匹配的记录
	//var users []UserInfo
	//db.Where("age > ?", 20).Find(&users)
	//for _, u := range users {
	//	fmt.Printf("获取所有匹配的记录[遍历]...%v \n", u)
	//}

	//根据条件查询
	//var users []UserInfo
	//db.Where("name <>  ?", "孙权").Find(&users)
	//for _, u := range users {
	//	fmt.Printf("根据条件查询[遍历]...%v \n", u)
	//}

	//测试 IN
	//var users []UserInfo
	//// 二选一----- db.Where("age IN (?,?)", 22, 23).Find(&users)
	//db.Where("name In (?)", []string{"王老五", "孙权"}).Find(&users)
	//for _, u := range users {
	//	fmt.Printf("测试 IN条件 [遍历]...%v \n", u)
	//}

	//测试 LIKE
	//var users []UserInfo
	//db.Where("name like ?", "%孙%").Find(&users)
	//for _, u := range users {
	//	fmt.Printf("测试 LIKE 条件 [遍历]...%v \n", u)
	//}

	//测试 AND
	//var users []UserInfo
	//db.Where("name = ? and age = ?", "赵六", 21).Find(&users)
	//for _, u := range users {
	//	fmt.Printf("测试 AND 条件 [遍历]...%v \n", u)
	//}

	//测试 Time
	//var users []UserInfo
	//db.Where("update_at", time.Now().Format("2006-01-02")).Find(&users)
	//for _, u := range users {
	//	fmt.Printf("测试 Time 条件 [遍历]...%v \n", u)
	//}

	//测试 BETWEEN
	//var users []UserInfo
	//today := time.Now()
	//lastweek := today.AddDate(0, 0, -7)
	//db.Where("created_at  BETWEEN ? AND ?", lastweek, today).Find(&users)
	//for _, u := range users {
	//	fmt.Printf("测试 BETWEEN 条件 [遍历]...%v \n", u)
	//}

	//测试 Struct
	//var users []UserInfo
	//db.Where(&UserInfo{Name: "王老五", Age: 20}).Find(&users)
	//for _, u := range users {
	//	fmt.Printf("测试 Struct 条件 [遍历]...%v \n", u)
	//}

	//测试 Map
	//var users []UserInfo
	//db.Where(map[string]interface{}{"name": "孙权", "age": 22}).Find(&users)
	//for _, u := range users {
	//	fmt.Printf("测试 Map 条件 [遍历]...%v \n", u)
	//}

	//测试 主键切片
	//var users []UserInfo
	//db.Where([]int{2, 3}).Find(&users)
	//for _, u := range users {
	//	fmt.Printf("测试 主键切片 条件 [遍历]...%v \n", u)
	//}

	//-----------------------------通过结构体测试零值-------------------start-----------------

	// 使用指针
	//type UserInfo struct {
	//	gorm.Model
	//	Name string
	//	Age  *int
	//}
	//// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	//db.SingularTable(true)
	//
	////自适应创建表
	//db.AutoMigrate(&UserInfo{})
	//
	////创建数据
	////db.Create(&UserInfo{Name: "晓琴"})
	////db.Create(&UserInfo{Name: "浪潮", Age: new(int)})
	//
	//var users []UserInfo
	////db.Where(&UserInfo{Name: "晓琴"}).Find(&users)
	//db.Where(&UserInfo{Name: "浪潮", Age: new(int)}).Find(&users)
	//for _, u := range users {
	//	fmt.Printf("测试 Struct 条件 [遍历]...%v \n", u)
	//}

	//-----------------------------通过结构体测试零值-------------------end-----------------
	//-----------------------------Scanner/Valuer-------------------start-----------------
	// 使用 Scanner/Valuer
	//type UserInfo struct {
	//	gorm.Model
	//	Name string
	//	Age  sql.NullInt64 // sql.NullInt64 实现了 Scanner/Valuer 接口
	//}
	//
	//// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	//db.SingularTable(true)
	//
	////自适应创建表
	//db.AutoMigrate(&UserInfo{})
	//
	//var users []UserInfo
	//db.Where(&UserInfo{Name: "浪潮", Age: sql.NullInt64{25, true}}).Find(&users)
	//for _, u := range users {
	//	fmt.Printf("测试 Struct 条件 [遍历]...%v \n", u)
	//}

	//-----------------------------Scanner/Valuer-------------------end-----------------

	//-----------------------------NOT-------------------start-----------------

	//测试 NOT
	//var users []UserInfo
	//db.Debug().Not("name", "赵六").Find(&users)
	//for _, u := range users {
	//	fmt.Printf("测试 NOT 条件 [遍历]...%v \n", u)
	//}

	//测试 NOT IN
	//var users []UserInfo
	//db.Not("name", []string{"王老五", "吴越"}).Find(&users)
	//for _, u := range users {
	//	fmt.Printf("测试 NOT IN 条件 [遍历]...%v \n", u)
	//}

	//测试 不在主键切片中的数据
	//var users []UserInfo
	//db.Not([]int{2, 3, 4}).Find(&users)
	//for _, u := range users {
	//	fmt.Printf("不在主键切片中的数据 [遍历]...%v \n", u)
	//}

	//获取第一条数据
	//var user UserInfo
	//db.Not([]int{}).First(&user)
	//fmt.Printf("获取第一条数据...%v \n", user)

	//常规SQL
	//var user UserInfo
	//db.Not("name = ?", "王老五").First(&user)
	//fmt.Printf("获取常规SQL数据...%v \n", user)

	//Struct
	//var user UserInfo
	//db.Not(&UserInfo{Name: "孙权"}).First(&user)
	//fmt.Printf("测试 Struct 条件...%v \n", user)

	//-----------------------------NOT-------------------end-----------------

	//Or 条件
	//var users []UserInfo
	//db.Where(&UserInfo{Name: "孙权"}).Or("age >= 22").First(&users)
	//for _, user := range users {
	//	fmt.Printf("测试 Or 条件...%v \n", user)
	//}

	//-------------------Inline Condition 内联条件---------begin-----------
	// 通过主键获取 (只适用于整数主键)
	//var users []UserInfo
	//db.First(&users, 2)
	//for _, user := range users {
	//	fmt.Printf("测试 主键获取 内联条件...%v \n", user)
	//}

	//如果是一个非整数类型，则通过主键获取
	//var users []UserInfo
	//db.First(&users, "id = ?", "string_primary_key")
	//for _, user := range users {
	//	fmt.Printf("测试 主键非整数类型 内联条件...%v \n", user)
	//}

	//Plain SQL
	//var users []UserInfo
	////方式1
	////db.Find(&users, "name like ?", "%王%")
	////方式2
	//db.Find(&users, "name <> ? and age <> ?", "王老五", 20)
	//for _, user := range users {
	//	fmt.Printf("测试 Plain SQL  条件[遍历]...%v \n", user)
	//}

	// Struct
	//var users []UserInfo
	//db.Find(&users, UserInfo{Age: 23})
	//for _, user := range users {
	//	fmt.Printf("测试 Struct  条件[遍历]...%v \n", user)
	//}

	//map
	//var users []UserInfo
	//db.Find(&users, map[string]interface{}{"name": "吴越"})
	//for _, user := range users {
	//	fmt.Printf("测试 map  条件[遍历]...%v \n", user)
	//}
	//-------------------Inline Condition 内联条件---------end-----------

	//-------------------Extra Querying option---------begin-----------

	//var users []UserInfo
	//db.Debug().Set("gorm:query_option", "FOR UPDATE").First(&users, 2)
	//for _, user := range users {
	//	fmt.Printf("测试 map  条件[遍历]...%v \n", user)
	//}

	//-------------------Extra Querying option---------end-----------

	//-------------------FirstOrInit---------begin-----------
	//获取匹配的第一条记录，否则根据给定的条件初始化一个新的对象 (仅支持 struct 和 map 条件)
	//var user UserInfo
	////找不到记录就初始化一条
	//db.Debug().FirstOrInit(&user, map[string]interface{}{"Name": "jinzhu", "Age": 26})
	//fmt.Printf("测试 匹配的第一条记录  条件[遍历]...%v \n", user)

	//找到数据-map
	//var user UserInfo
	//db.FirstOrInit(&user, map[string]interface{}{"Name": "王老五"})
	//fmt.Printf("测试 FirstOrInit 找到数据  条件[遍历]...%v \n", user)

	//找到数据-struct
	//var user UserInfo
	//db.Where(UserInfo{Name: "老七", Age: 27}).FirstOrInit(&user)
	//fmt.Printf("测试 FirstOrInit 找到数据-struct  条件...%v \n", user)
	//-------------------FirstOrInit---------end-----------

	//-------------------FirstOrInit.Attrs---------begin-----------
	//未找到-struct
	//var user UserInfo
	//db.Where(UserInfo{Name: "老八"}).Attrs(UserInfo{Age: 15}).FirstOrInit(&user)
	//fmt.Printf("测试 Attrs 未找到数据-struct  条件...%v \n", user)

	//未找到
	//var user UserInfo
	//db.Where(UserInfo{Name: "老八"}).Attrs("age", 15).FirstOrInit(&user)
	//fmt.Printf("测试 Attrs 未找到数据-属性替换  ...%v \n", user)

	//找到数据-保留原数据
	//var user UserInfo
	//db.Where(UserInfo{Name: "王老五"}).Attrs("age", 15).FirstOrInit(&user)
	//fmt.Printf("测试 Attrs 未找到数据-属性替换  ...%v \n", user)

	//-------------------FirstOrInit.Attrs---------end-----------

	//-------------------FirstOrInit.Assign---------begin-----------
	//不管记录是否找到，都将参数赋值给 struct.

	//var user UserInfo
	//// 未找到
	//db.Where(UserInfo{Name: "老八"}).Assign("age", 15).FirstOrInit(&user)
	//fmt.Printf("测试 Assign 未找到数据-属性替换  ...%v \n", user)

	//已找到-属性替换
	//var user UserInfo
	//db.Where(UserInfo{Name: "王老五"}).Assign("age", 15).FirstOrInit(&user)
	//fmt.Printf("测试 Assign 找到数据-属性替换  ...%v \n", user)

	//已找到-struct
	//var user UserInfo
	//db.Where(UserInfo{Name: "王老五"}).Assign(UserInfo{Age: 15}).FirstOrInit(&user)
	//fmt.Printf("测试 Assign 找到数据-struct  ...%v \n", user)
	//-------------------FirstOrInit.Assign---------end-----------

	//-------------------FirstOrCreate---------begin-----------

	// 未找到
	//var user UserInfo
	//db.FirstOrCreate(&user, UserInfo{Name: "non_existing"})
	//fmt.Printf("测试 FirstOrCreate 未找到数据  ...%v \n", user)

	// 已找到
	//var user UserInfo
	//db.FirstOrCreate(&user, UserInfo{Name: "王老五"})
	//fmt.Printf("测试 FirstOrCreate 已找到数据  ...%v \n", user)

	// Attrs- 如果记录未找到，将使用参数创建 struct 和记录.
	//var user UserInfo
	////未找到
	//db.Where(UserInfo{Name: "non_existing3"}).Attrs(UserInfo{Age: 15}).FirstOrCreate(&user)
	//fmt.Printf("测试 FirstOrCreate - Attrs未找到数据  ...%v \n", user)

	// 已找到
	//var user UserInfo
	//db.Where(UserInfo{Name: "赵六"}).Attrs(UserInfo{Age: 15}).FirstOrCreate(&user)
	//fmt.Printf("测试 FirstOrCreate - Attrs已找到数据  ...%v \n", user)

	//Assign-不管记录是否找到，都将参数赋值给 struct 并保存至数据库.
	//var user UserInfo
	////未找到
	//db.Where(UserInfo{Name: "non_existing4"}).Assign(UserInfo{Age: 15}).FirstOrCreate(&user)
	//fmt.Printf("测试 FirstOrCreate - Assign未找到数据  ...%v \n", user)

	//已找到
	//var user UserInfo
	//db.Debug().Where(UserInfo{Name: "孙权"}).Assign(UserInfo{Age: 15}).FirstOrCreate(&user)
	//fmt.Printf("测试 FirstOrCreate - Assign找到数据  ...%v \n", user)

	//-------------------FirstOrCreate---------end-----------

	//-------------------Advanced Query---------begin-----------
	//SubQuery 子查询
	//var users []UserInfo
	////生成的SQL为:  SELECT * FROM `user_info`  WHERE `user_info`.`deleted_at` IS NULL AND ((age >= (SELECT AVG(age) FROM `user_info`  WHERE (name like '%吴%'))))
	//db.Debug().Where("age >= ?", db.Table("user_info").Select("AVG(age)").Where("name like ?", "%吴%").SubQuery()).Find(&users)
	//for _, user := range users {
	//	fmt.Printf("测试 子查询 数据[遍历]  ...%v \n", user)
	//}

	//选择字段-1
	//var users []UserInfo
	//db.Debug().Select("name,age").Find(&users)
	//for _, user := range users {
	//	fmt.Printf("测试 选择字段-1 数据[遍历]  ...%v \n", user)
	//}

	//选择字段-2
	//var users []UserInfo
	//db.Debug().Select([]string{"name", "age"}).Find(&users)
	//for _, user := range users {
	//	fmt.Printf("测试 选择字段-2 数据[遍历]  ...%v \n", user)
	//}

	//选择字段-3------COALESCE用法没搞懂
	//var users []UserInfo
	//db.Debug().Table("user_info").Select("COALESCE(age,?)", 21).Rows()
	//for _, user := range users {
	//	fmt.Printf("测试 选择字段-3 数据[遍历]  ...%v \n", user)
	//}

	//Order-1
	//var users []UserInfo
	//db.Debug().Order("age desc , name").Find(&users)
	//for _, user := range users {
	//	fmt.Printf("测试 Order-1 数据[遍历]  ...%v \n", user)
	//}

	//Order-2   多字段排序
	//var users []UserInfo
	//db.Debug().Order("age desc").Order("name").Find(&users)
	//for _, user := range users {
	//	fmt.Printf("测试 Order-2 数据[遍历]  ...%v \n", user)
	//}

	//Order-3   覆盖排序
	//var users1 []UserInfo
	//var users2 []UserInfo
	//db.Debug().Order("age desc").Find(&users1).Order("age", true).Find(&users2)
	//for index, user := range users1 {
	//	fmt.Printf("测试 Order-3-users1 数据[遍历]  %d...%v \n", index+1, user)
	//}
	//for idx, user := range users2 {
	//	fmt.Printf("测试 Order-3-users2 数据[遍历]  %d...%v \n", idx+1, user)
	//}

	//Limit-1   指定从数据库检索出的最大记录数。
	//var users []UserInfo
	//db.Debug().Limit(3).Find(&users)
	//for _, user := range users {
	//	fmt.Printf("测试 Limit-1 数据[遍历]  ...%v \n", user)
	//}

	//Limit-2   -1 取消 Limit 条件
	//var users1 []UserInfo
	//var users2 []UserInfo
	//db.Debug().Limit(5).Find(&users1).Limit(-1).Find(&users2)
	//for index, user := range users1 {
	//	fmt.Printf("测试 Order-2-users1 数据[遍历]  %d...%v \n", index+1, user)
	//}
	//for idx, user := range users2 {
	//	fmt.Printf("测试 Order-2-users2 数据[遍历]  %d...%v \n", idx+1, user)
	//}

	//Offset-1   指定开始返回记录前要跳过的记录数。
	//var users []UserInfo
	//db.Debug().Limit(4).Offset(8).Find(&users)
	//for _, user := range users {
	//	fmt.Printf("测试 Offset-1 数据[遍历]  ...%v \n", user)
	//}

	//Offset-2   -1 取消 Limit 条件
	//var users1 []UserInfo
	//var users2 []UserInfo
	//db.Debug().Limit(2).Offset(5).Find(&users1).Offset(-1).Find(&users2)
	//for index, user := range users1 {
	//	fmt.Printf("测试 Offset-2-users1 数据[遍历]  %d...%v \n", index+1, user)
	//}
	//for idx, user := range users2 {
	//	fmt.Printf("测试 Offset-2-users2 数据[遍历]  %d...%v \n", idx+1, user)
	//}

	//Count-1   该 model 能获取的记录总数。
	//var users []UserInfo
	//var count int
	//db.Debug().Where("name = ?", "吴越").Or("name like ?", "%王%").Find(&users).Count(&count)
	//fmt.Printf("总共有 %d 条数据 \n", count)
	//for _, user := range users {
	//	fmt.Printf("测试 Count-1 数据[遍历]  ...%v \n", user)
	//}

	//Count-2
	//var users []UserInfo
	//var count int
	//db.Debug().Model(&UserInfo{}).Where("name like ?", "%吴%").Find(&users).Count(&count)
	//fmt.Printf("总共有 %d 条数据 \n", count)
	//for _, user := range users {
	//	fmt.Printf("测试 Count-1 数据[遍历]  ...%v \n", user)
	//}

	//Count-3
	//var count int
	//db.Debug().Table("user_info").Count(&count)
	//fmt.Printf("总共有 %d 条数据 \n", count)

	//Count-4
	//var count int
	//db.Debug().Table("user_info").Select("count(distinct(age))").Count(&count)
	//fmt.Printf("总共有 %d 条数据 \n", count)

	//Group & Having
	//rows, err := db.Table("user_info").Select("date(create_at) as date , sum(age) as total").Group("date(create_at)").Rows()
	////遍历报错
	//for rows.Next() {
	//	columns, _ := rows.Columns()
	//	fmt.Printf("测试 Count-1 数据[遍历]  ...%v \n", columns)
	//}

	//Pluck-1   查询 model 中的一个列作为切片，如果您想要查询多个列，您应该使用 Scan
	//var users []UserInfo
	//var ages []int
	//db.Debug().Find(&users).Pluck("age", &ages)
	//for _, user := range users {
	//	fmt.Printf("测试 Pluck-1 数据[遍历]  ...%v \n", user)
	//}
	//for inx, age := range ages {
	//	fmt.Printf("年龄 数据[遍历]  %d...%d \n", inx, age)
	//}

	//Pluck-2   查询 model 中的一个列作为切片，如果您想要查询多个列，您应该使用 Scan
	//var names []string
	//db.Debug().Model(&UserInfo{}).Pluck("name", &names)
	//for inx, name := range names {
	//	fmt.Printf("Pluck-2 姓名 数据[遍历]  %d...%s \n", inx, name)
	//}

	//Pluck-3
	//var names []string
	//db.Debug().Table("user_info").Pluck("name", &names)
	//for inx, name := range names {
	//	fmt.Printf("Pluck-3 姓名 数据[遍历]  %d...%s \n", inx, name)
	//}

	//查询多个字段
	//var users []UserInfo
	//db.Debug().Select("name,age").Find(&users)
	//for inx, user := range users {
	//	fmt.Printf("查询多个字段 数据[遍历]  %d...%v \n", inx, user)
	//}

	//Scan，扫描结果至一个 struct.
	type Result struct {
		Name string
		Age  int
	}

	var results []Result
	//方式1
	//db.Table("user_info").Select("name,age").Where("name like ?", "%吴%").Scan(&results)
	//方式2-原生 SQL
	db.Raw("select name , age from user_info where name like ?", "%吴%").Scan(&results)
	for _, result := range results {
		fmt.Printf("测试 SCAN 数据[遍历]  ...%v \n", result)
	}
}
