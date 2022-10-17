package logic

import (
	"goweb32_bells-of-ireland/dao/mysql"
	"goweb32_bells-of-ireland/models"
	"goweb32_bells-of-ireland/pkg/snowflake"
)

func CreateArticle(p *models.Post) error {
	//1. 生成ID
	p.ID = snowflake.GenID()
	//2. 插入数据
	return mysql.InsertPost(p)
}

func QueryArticleDetail(id int64) (*models.Post, error) {

	return mysql.QueryPostById(id)
}
