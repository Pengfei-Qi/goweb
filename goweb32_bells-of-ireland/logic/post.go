package logic

import (
	"goweb32_bells-of-ireland/dao/mysql"
	"goweb32_bells-of-ireland/dao/redis"
	"goweb32_bells-of-ireland/models"
	"goweb32_bells-of-ireland/pkg/snowflake"

	"go.uber.org/zap"
)

func CreateArticle(p *models.Post) error {
	//1. 生成ID
	p.ID = snowflake.GenID()
	//2. 插入数据
	if err := mysql.InsertPost(p); err != nil {
		return err
	}
	if err := redis.CreatePostTime(p); err != nil {
		return err
	}

	return nil
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

func QueryPostListDetail2(param *models.ParamPostData) (apiPosts []*models.ApiPostDetail, err error) {

	//查询redis中的postID
	ids, err := redis.GetPostIDsByPage(param)
	if err != nil {
		zap.L().Error("redis.GetPostIDsByPage is err")
		return
	}
	if len(ids) == 0 {
		zap.L().Error("redis get pods is null")
		return
	}
	zap.L().Info("redis query posts is", zap.Any("posts", ids))
	//根据postID 查询数据库中的 posts数据
	posts, err := mysql.QueryPostsByIDs(ids)
	if err != nil {
		zap.L().Error("post data not found ",
			zap.Error(err))
		return
	}
	voteData, err := redis.GetPostsVoteNum(ids)
	if err != nil {
		zap.L().Error("redis.GetPostsVoteNum error", zap.Error(err))
		return
	}
	zap.L().Info("voteData data is: ", zap.Any("vote", voteData))
	apiPosts = make([]*models.ApiPostDetail, 0, len(posts))
	for inx, post := range posts {
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
			VoteNum:         voteData[inx],
			Post:            post,
			CommunityDetail: community,
		}

		apiPosts = append(apiPosts, data)
	}

	return
}

func QueryCommunityPostListDetail(param *models.ParamPostData) (apiPosts []*models.ApiPostDetail, err error) {
	//查询redis中的postID
	ids, err := redis.GetCommunityPostIDsByOrder(param)
	if err != nil {
		zap.L().Error("redis.GetCommunityPostIDsByOrder is err")
		return
	}
	if len(ids) == 0 {
		zap.L().Error("redis get posts is null")
		return
	}
	zap.L().Info("redis query community posts is", zap.Any("posts", ids))
	//根据postID 查询数据库中的 posts数据
	posts, err := mysql.QueryPostsByIDs(ids)
	if err != nil {
		zap.L().Error("community post data not found ",
			zap.Error(err))
		return
	}
	voteData, err := redis.GetPostsVoteNum(ids)
	if err != nil {
		zap.L().Error("redis.GetPostsVoteNum error", zap.Error(err))
		return
	}
	zap.L().Info("community voteData data is: ", zap.Any("vote", voteData))
	apiPosts = make([]*models.ApiPostDetail, 0, len(posts))
	for inx, post := range posts {
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
			VoteNum:         voteData[inx],
			Post:            post,
			CommunityDetail: community,
		}

		apiPosts = append(apiPosts, data)
	}

	return
}

func QueryPostListNew(param *models.ParamPostData) (apiPosts []*models.ApiPostDetail, err error) {
	if param.CommunityID == 0 {
		apiPosts, err = QueryPostListDetail2(param)
	} else {
		apiPosts, err = QueryCommunityPostListDetail(param)
	}
	if err != nil {
		zap.L().Error("QueryPostListNew get data failed", zap.Error(err))
		return nil, err
	}
	return
}
