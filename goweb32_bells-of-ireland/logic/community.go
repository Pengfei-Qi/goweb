package logic

import (
	"goweb32_bells-of-ireland/dao/mysql"
	"goweb32_bells-of-ireland/models"
)

func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (communityDetail *models.CommunityDetail, err error) {
	return mysql.GetCommunityDetailByID(id)
}
