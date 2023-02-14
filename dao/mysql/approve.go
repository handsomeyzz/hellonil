package mysql

import (
	"go.uber.org/zap"
	"hellonil/dao/redis"
	"hellonil/pkg/snowflake"
)

// InsertLike 插入喜欢记录
func InsertLike(userId int64, videoId string) error {
	id := snowflake.GenID()
	sqlStr := `insert into likes values (?, ?, ?, 0) `
	_, err := db.Exec(sqlStr, id, userId, videoId)
	if err != nil {
		return err
	}
	return nil
}

// 删除喜欢记录
func DeleteLike(userId int64, videoId string) error {
	sqlstr := `delete from likes where user_id = ? and video_id = ?`
	_, err := db.Exec(sqlstr, userId, videoId)
	if err != nil {
		return err
	}
	return nil
}

// InquireFavorite 查询视频点赞数 利用redis加缓存
func InquireFavorite(videoId string) (n int, err error) {
	n, err = redis.InquireFavorite(videoId)
	if err == nil {
		//查询redis成功，不查MySQL
		return n, nil
	}
	sqlStr := `select favorite_count from videos where id = ?`
	row := db.QueryRow(sqlStr, videoId)
	row.Scan(&n)
	if n != -1 {
		//根据MySQL的值更新redis
		err := redis.UpdateFavorite(videoId, n)
		if err != nil {
			zap.L().Info("update redis failed")
		}
		return n, nil
	}
	return n, err
}

// UpdateFavorite 更新点赞数
func UpdateFavorite(videoId string, n int) error {
	sqlStr := `update videos set favorite_count = ? where id = ?`
	_, err := db.Exec(sqlStr, n, videoId)
	if err != nil {
		return err
	}
	//同时更新redis
	err2 := redis.UpdateFavorite(videoId, n)
	if err2 != nil {
		zap.L().Info("update redis failed")
	}
	return nil
}
