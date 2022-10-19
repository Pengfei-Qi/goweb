package logic

import (
	"goweb32_bells-of-ireland/dao/mysql"
	"goweb32_bells-of-ireland/models"
	"goweb32_bells-of-ireland/pkg/snowflake"

	"go.uber.org/zap"
)

func CreateArticle(p *models.Post) error {
	//1. 生成ID
	p.ID = snowflake.GenID()
	//2. 插入数据
	return mysql.InsertPost(p)
}

func QueryArticleDetail(id int64) (data *models.ApiPostDetail, err error) {
	//查询post
	post, err := mysql.QueryPostById(id)
	if err != nil {
		zap.L().Error("post data not found ",
			zap.Int64("post_id", id),
			zap.Error(err))
		return
	}

	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("user data not found ",
			zap.Int64("user_id", post.AuthorID),
			zap.Error(err))
		return
	}

	//获取community
	community, err := mysql.GetCommunityDetailByID(post.CommunityId)
	if err != nil {
		zap.L().Error("community data not found ",
			zap.Int64("community_data", post.CommunityId),
			zap.Error(err))
		return
	}
	//封装数据
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}

	return data, err
}

func QueryPostListDetail(page, size int64) (apiPosts []*models.ApiPostDetail, err error) {
	//查询post
	posts, err := mysql.QueryPostList(page, size)
	if err != nil {
		zap.L().Error("post data not found ",
			zap.Error(err))
		return
	}

	apiPosts = make([]*models.ApiPostDetail, 0, len(posts))

	for _, post := range posts {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("user data not found ",
				zap.Int64("user_id", post.AuthorID),
				zap.Error(err))
			continue
		}

		//获取community
		community, err := mysql.GetCommunityDetailByID(post.CommunityId)
		if err != nil {
			zap.L().Error("community data not found ",
				zap.Int64("community_data", post.CommunityId),
				zap.Error(err))
			continue
		}
		//封装数据
		data := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}

		apiPosts = append(apiPosts, data)
	}

	return
}
