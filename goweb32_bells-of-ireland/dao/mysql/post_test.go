package mysql

import (
	"fmt"
	"goweb32_bells-of-ireland/models"
	"testing"

	"github.com/jmoiron/sqlx"
)

func init() {
	// DSN:Data Source Name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		"root",
		"123456",
		"192.168.235.233",
		3306,
		"bells_of_ireland")
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, _ = sqlx.Connect("mysql", dsn)
}

func TestInsertPost(t *testing.T) {
	postData := &models.Post{
		ID:          12345,
		AuthorID:    2333,
		CommunityId: 2,
		Title:       "test",
		Content:     "三生三世十里桃花",
	}

	err := InsertPost(postData)
	if err != nil {
		t.Fatal("TestInsertPost insert post failed")
	}
	t.Logf("TestInsertPost insert post successed")
}
