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
	p.post_id,
	p.title,
	p.content,
	p.community_id,
	c.community_name 
	FROM
		post p,
		community c 
	WHERE
	 p.community_id = c.community_id AND
		p.post_id = ?`
	err = db.Get(post, sqlStr, postID)
	return
}
