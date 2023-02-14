package mysql

import (
	"go.uber.org/zap"
	"hellonil/models"
	"hellonil/pkg/snowflake"
)

func InsertVideo(video *models.Videos) (err error) {
	sqlStr := `insert into videos(id,author_id,play_url,cover_url,favorite_count,comment_count,title,create_time) VALUES (?,?,?,?,?,?,?,?)`
	id := snowflake.GenID()
	_, err = db.Exec(sqlStr, id, video.AuthorID, video.PlayUrl, video.CoverUrl, video.FavoriteCount, video.CommentCount, video.Title, video.CreateTime)
	if err != nil {
		zap.L().Info("插入视频数据失败", zap.Error(err))
		return err
	}
	return nil
}

func CheckVidExist(vid int64) bool {
	sqlstr := `select id from videos where id = ?`
	var v int64
	err := db.Get(&v, sqlstr, vid)
	if err == nil {
		return true
	}
	return false
}
