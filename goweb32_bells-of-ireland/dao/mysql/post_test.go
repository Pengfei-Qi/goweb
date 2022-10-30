package mysql

import (
	"fmt"
	"goweb32_bells-of-ireland/models"
	"goweb32_bells-of-ireland/settings"
	"testing"
)

func init() {
	abc := &settings.MysqlConfig{
		Host:         "192.168.235.233",
		Username:     "root",
		Password:     "123456",
		DbName:       "bells_of_ireland",
		Port:         3306,
		MaxOpenConns: 100,
		MaxIdleConns: 20,
	}
	err := Init(abc)
	if err != nil {
		fmt.Printf("mysql connect failed , because of :%s", err.Error())
	}
	fmt.Printf("mysql connect success")
}

func TestInsertPost(t *testing.T) {
	postData := &models.Post{
		ID:          5632,
		AuthorID:    2222,
		CommunityId: 1,
		Title:       "test3",
		Content:     "黄林",
	}

	err := InsertPost(postData)
	if err != nil {
		t.Fatal("TestInsertPost insert post failed")
	}
	t.Logf("TestInsertPost insert post successed")
}
