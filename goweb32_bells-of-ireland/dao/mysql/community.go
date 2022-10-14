package mysql

import (
	"database/sql"
	"fmt"
	"goweb32_bells-of-ireland/models"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id,community_name from community"
	err = db.Select(&communityList, sqlStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("there is no data")
		}
	}
	return
}
