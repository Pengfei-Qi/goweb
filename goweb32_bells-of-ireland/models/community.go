package models

import "time"

type Community struct {
	ID   int64  `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

type CommunityDetail struct {
	ID           int64     `json:"id" db:"community_id"`                     //社区id
	Name         string    `json:"name" db:"community_name"`                 //社区名称
	Introduction string    `json:"introduction,omitempty" db:"introduction"` //介绍
	CreateTime   time.Time `json:"createTime" db:"create_time"`              //创建时间
}
