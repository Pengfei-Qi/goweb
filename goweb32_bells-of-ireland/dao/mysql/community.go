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

func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {

	community = new(models.CommunityDetail)

	sqlStr := `SELECT
	community_id,
	community_name,
	introduction,
	create_time
	FROM
		community
	WHERE
		community_id = ?`
	if err = db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorPramValid
		}
	}
	return community, err
}
