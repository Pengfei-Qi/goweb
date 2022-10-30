package models

import "time"

type Post struct {
	ID          int64     `json:"id,string" db:"post_id"`                            //帖子id
	AuthorID    int64     `json:"author_id,string" db:"author_id"`                   //作者id
	CommunityId int64     `json:"community_id" db:"community_id" binding:"required"` //社区ID
	Status      int32     `json:"status" db:"status"`                                //状态
	Title       string    `json:"title" db:"title" binding:"required"`               //标题
	Content     string    `json:"content" db:"content" binding:"required"`           //内容
	CreateTime  time.Time `json:"create_time" db:"create_time"`                      //创建时间
}

type ApiPostDetail struct {
	AuthorName string `json:"author_name"` //用户名称
	VoteNum    int64  `json:"vote_num"`    //投票数量
	*Post
	*CommunityDetail `json:"community"`
}
