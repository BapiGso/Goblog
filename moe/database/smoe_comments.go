package database

// CommentsWithCid 根据文章cid查询该文章的评论
func (s *S) CommentsWithCid(cid int) error {
	err := db.Select(&s.CommArr, `SELECT * FROM  typecho_comments 
		WHERE cid=?
		ORDER BY created`, cid)
	return err
}
