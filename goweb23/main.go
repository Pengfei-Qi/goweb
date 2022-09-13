package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(192.168.235.233:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	DB.SetMaxOpenConns(20)
	DB.SetConnMaxIdleTime(10)
	return
}

type User struct {
	ID   int    `db:"id"`
	Age  int    `db:"age"`
	Name string `db:"name"`
}

// 查询单条数据示例
func queryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var u User
	err := DB.Get(&u, sqlStr, 2)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
}

// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	var users []User
	err := DB.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	for _, user := range users {
		fmt.Printf("users:%#v\n", user)
	}
}

// 插入数据
func insertRowDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := DB.Exec(sqlStr, "沙河小王子", 19)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRowDemo() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := DB.Exec(sqlStr, 39, 6)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := DB.Exec(sqlStr, 6)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

func insertUserDemo() (err error) {
	sqlStr := "INSERT INTO user (name,age) VALUES (:name,:age)"
	_, err = DB.NamedExec(sqlStr,
		map[string]interface{}{
			"name": "大窑",
			"age":  35,
		})
	return
}
func main() {
	err := initDB()
	if err != nil {
		panic(err)
		return
	}
	defer DB.Close()
	//queryRowDemo()
	//queryMultiRowDemo()
	insertUserDemo()
	fmt.Printf("执行完成")
}
