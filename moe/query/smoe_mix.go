package query

import "github.com/jmoiron/sqlx"

//join查询的sql都放这里

type CommsWithTitleMix struct {
	Comments
	Title string `db:"title"     form:"title"`
}

// CommsWithTitle 查询评论组，后台专用
func CommsWithTitle(Db *sqlx.DB, status string, limit, pagenum int) []CommsWithTitleMix {
	data := make([]CommsWithTitleMix, 0, limit)
	_ = Db.Select(&data, `SELECT c.*,title
    	FROM typecho_comments AS c 
        INNER JOIN typecho_contents on typecho_contents.cid=c.cid
		WHERE c.status=? 
		ORDER BY c.created DESC 
		LIMIT ? OFFSET ?`, status, limit, pagenum*limit-limit)
	return data
}

type PostsWithCoverMix struct {
	Contents
	Cover string `db:"cover"     form:"cover"`
}

// PostsWithCover 查询带封面的文章组
func PostsWithCover(Db *sqlx.DB, status string, limit, pagenum int) []PostsWithCoverMix {
	data := make([]PostsWithCoverMix, 0, limit)
	_ = Db.Select(&data, `SELECT c.*,f.str_value AS cover
    	FROM typecho_contents AS c 
        INNER JOIN typecho_fields AS f on f.cid=c.cid
		WHERE status=? AND f.name='coverList'
		ORDER BY c.cid DESC 
		LIMIT ? OFFSET ?`, status, limit, pagenum*limit-limit)
	return data
}

type PostWithBGMusicMix struct {
	Contents
	Music string `db:"music"     form:"music"`
}

// PostWithBGMusic 查询带封面的文章
func PostWithBGMusic(Db *sqlx.DB, status string, limit, pagenum int) PostWithBGMusicMix {
	data := PostWithBGMusicMix{}
	_ = Db.Get(&data, `SELECT c.*,f.str_value AS music
    	FROM typecho_contents AS c 
        INNER JOIN typecho_fields AS f on f.cid=c.cid
		WHERE status=? AND f.name='bgMusicList'
		ORDER BY c.cid DESC 
		LIMIT ? OFFSET ?`, status, limit, pagenum*limit-limit)
	return data
}

type MediasWithTitleMix struct {
	Contents
	PostTitle string `db:"post_title"     form:"post_title"`
}

func MediasWithTitle(Db *sqlx.DB, limit, pagenum int) []MediasWithTitleMix {
	data := make([]MediasWithTitleMix, 0, limit)
	_ = Db.Select(&data, `SELECT c1.*,c2.title AS post_title
    	FROM typecho_contents AS c1 
        INNER JOIN typecho_contents AS c2 on c1.parent=c2.cid
		ORDER BY c1.cid DESC 
		LIMIT ? OFFSET ?`, limit, pagenum*limit-limit)
	return data
}
