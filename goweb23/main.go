package main

import (
	"database/sql/driver"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
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

func (u User) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
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
			"name": "可口可乐",
			"age":  31,
		})
	return
}

func namedQuery() {
	sqlStr := "SELECT * FROM user WHERE age=:age"
	// 使用map做命名查询
	rows, err := DB.NamedQuery(sqlStr, map[string]interface{}{"age": 31})
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u User
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user1:%#v\n", u)
	}

	u := User{
		Age: 31,
	}
	// 使用结构体命名查询，根据结构体字段的 db tag进行映射
	rows, err = DB.NamedQuery(sqlStr, u)
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u User
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user2:%#v\n", u)
	}
}

// 事务操作
func transactionDemo2() (err error) {
	tx, err := DB.Beginx() // 开启事务
	if err != nil {
		fmt.Printf("begin trans failed, err:%v\n", err)
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			fmt.Println("rollback")
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
			fmt.Println("commit")
		}
	}()

	sqlStr1 := "Update user set age=20 where id = ?"

	rs, err := tx.Exec(sqlStr1, 1)
	if err != nil {
		return err
	}
	n, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		fmt.Printf("exec sqlStr1 failed, affectRowCount is : %d \n", n)
		return errors.New("exec sqlStr1 failed")
	}

	sqlStr2 := "update user set age=50 where id=?"
	rs, err = tx.Exec(sqlStr2, 4)
	if err != nil {
		return err
	}
	n, err = rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		fmt.Printf("exec sqlStr2 failed, affectRowCount is : %d \n", n)
		return errors.New("exec sqlStr2 failed")
	}
	return err
}

// BatchInsertUsers 自行构造批量插入的语句
func BatchInsertUsers(users []*User) error {
	// 存放 (?, ?) 的slice
	valueStrings := make([]string, 0, len(users))
	// 存放values的slice
	valueArgs := make([]interface{}, 0, len(users)*2)
	// 遍历users准备相关数据
	for _, u := range users {
		// 此处占位符要与插入值的个数对应
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, u.Name)
		valueArgs = append(valueArgs, u.Age)
	}
	// 自行拼接要执行的具体语句
	stmt := fmt.Sprintf("INSERT INTO user (name, age) VALUES %s",
		strings.Join(valueStrings, ","))
	_, err := DB.Exec(stmt, valueArgs...)
	return err
}

// BatchInsertUsers2 使用sqlx.In帮我们拼接语句和参数, 注意传入的参数是[]interface{}
func BatchInsertUsers2(users []interface{}) error {
	query, args, _ := sqlx.In(
		"INSERT INTO user (name, age) VALUES (?), (?), (?)",
		users..., // 如果arg实现了 driver.Valuer, sqlx.In 会通过调用 Value()来展开它
	)
	fmt.Println(query) // 查看生成的querystring
	fmt.Println(args)  // 查看生成的args
	_, err := DB.Exec(query, args...)
	return err
}

// BatchInsertUsers3 使用NamedExec实现批量插入
func BatchInsertUsers3(users []*User) error {
	_, err := DB.NamedExec("INSERT INTO user (name, age) VALUES (:name, :age)", users)
	return err
}

// QueryByIDs 根据给定ID查询
func QueryByIDs(ids []int) (users []User, err error) {
	query, args, err := sqlx.In("SELECT * FROM user WHERE id IN (?);", ids)
	if err != nil {
		fmt.Printf("查询失败,err:%v \n", err)
		return
	}
	query = DB.Rebind(query)
	err = DB.Select(&users, query, args...)
	return
}

// QueryAndOrderByIDs 按照指定id查询并维护顺序
func QueryAndOrderByIDs(ids []int) (users []User, err error) {
	// 动态填充id
	idsStr := make([]string, 0, len(ids))
	for _, id := range ids {
		idsStr = append(idsStr, fmt.Sprintf("%d", id))
	}
	query, args, err := sqlx.In("SELECT * FROM user WHERE id IN (?) ORDER BY FIND_IN_SET(id, ?);", ids, strings.Join(idsStr, ","))
	if err != nil {
		fmt.Printf("查询失败,err:%v \n", err)
		return
	}
	query = DB.Rebind(query)
	err = DB.Select(&users, query, args...)
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
	//insertUserDemo()
	//namedQuery()
	//transactionDemo2()

	//1. 批量插入-自己写的
	//user1 := User{Name: "小茗同学", Age: 40}
	//user2 := User{Name: "雀巢咖啡", Age: 41}
	//user3 := User{Name: "星巴克", Age: 42}
	//users := []*User{&user1,&user2,&user3}
	//err = BatchInsertUsers(users)
	//if err != nil {
	//	fmt.Printf("批量插入数据失败, err:%#v,\n", err)
	//	return
	//}

	//2. 批量插入-sqlIN
	//user1 := User{Name: "小茗同学_sqlIn", Age: 40}
	//user2 := User{Name: "雀巢咖啡_sqlIn", Age: 41}
	//user3 := User{Name: "星巴克_sqlIn", Age: 42}
	//err = BatchInsertUsers2([]interface{}{user1, user2, user3})
	//if err != nil {
	//	fmt.Printf("sql.in 批量插入数据失败, err:%#v,\n", err)
	//	return
	//}

	//3. 批量插入-NamedExec
	//user1 := User{Name: "小茗同学_NamedExec", Age: 45}
	//user2 := User{Name: "雀巢咖啡_NamedExec", Age: 45}
	//user3 := User{Name: "星巴克_NamedExec", Age: 45}
	//users := []*User{&user1, &user2, &user3}
	//err = BatchInsertUsers3(users)
	//if err != nil {
	//	fmt.Printf("sql.in 批量插入数据失败, err:%#v,\n", err)
	//	return
	//}

	//1. 批量查询-按照mysql的id默认排序方式
	//users, err := QueryByIDs([]int{15, 1, 16, 8})
	//if err != nil {
	//	fmt.Printf("sql.in 批量查询数据失败, err:%#v,\n", err)
	//	return
	//}
	//for _, user := range users {
	//	fmt.Printf("sql.in 批量查询到的数据为, user :%#v \n", user)
	//}

	//2. 批量查询-按照自定义排序方式
	users, err := QueryAndOrderByIDs([]int{15, 1, 16, 8})
	if err != nil {
		fmt.Printf("sql.in 批量查询数据失败, err:%#v,\n", err)
		return
	}
	for _, user := range users {
		fmt.Printf("sql.in 批量查询到的数据为, user :%#v \n", user)
	}
	fmt.Printf("执行完成")
}
