package mysql

import (
	"goweb32_bells-of-ireland/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

//sqlx例子: https://pkg.go.dev/github.com/jmoiron/sqlx@v1.3.5#section-readme

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

func QueryPostsByIDs(ids []string) (posts []*models.Post, err error) {
	//find_in_set用法: https://doris.apache.org/zh-CN/docs/dev/sql-manual/sql-functions/string-functions/find_in_set/
	posts = make([]*models.Post, len(ids))
	sqlStr := `
	SELECT DISTINCT
		post_id,
		title,
		content,
		community_id,
		author_id
	FROM
		post
	Where
		post_id in (?)
	ORDER BY 
		FIND_IN_SET(post_id, ?);`
	// 参考1: http://jmoiron.github.io/sqlx/
	// 参考2: FIND_IN_SET:     https://www.cnblogs.com/guyouyin123/p/14481196.html
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	query = db.Rebind(query)
	err = db.Select(&posts, query, args...)
	return
}
