package mysql

import (
	"fmt"
	"goweb32_bells-of-ireland/settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(mysqlConf *settings.MysqlConfig) (err error) {
	// DSN:Data Source Name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		mysqlConf.Username,
		mysqlConf.Password,
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.DbName)
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(mysqlConf.MaxOpenConns)
	db.SetMaxIdleConns(mysqlConf.MaxIdleConns)
	return

}
func Close() {
	_ = db.Close()
}
