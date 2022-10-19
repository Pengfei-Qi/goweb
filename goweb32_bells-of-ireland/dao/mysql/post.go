package mysql

import (
	"goweb32_bells-of-ireland/models"
)

func InsertPost(p *models.Post) (err error) {
	sqlStr := `INSERT INTO post (post_id, title, content,author_id, community_id) values (?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityId)
	return
}

func QueryPostById(postID int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `SELECT DISTINCT
	post_id,
	title,
	content,
	community_id,
    author_id
	FROM
		post 
	WHERE
	 post_id = ?`
	err = db.Get(post, sqlStr, postID)
	return
}

func QueryPostList(page, size int64) (posts []*models.Post, err error) {
	posts = make([]*models.Post, 2)
	sqlStr := `SELECT DISTINCT
	post_id,
	title,
	content,
	community_id,
    author_id
	FROM
		post
		limit ?,?`
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}
