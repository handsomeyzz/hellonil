package mysql

import (
	"go.uber.org/zap"
	"hellonil/models"
	"hellonil/responseStruct"
)

func LikeList(user_id string) (res []*responseStruct.Video, err error) {
	sqlstr0 := `select id,author_id,play_url,cover_url,favorite_count,comment_count,title from videos where  id = any(select video_id from likes WHERE user_id = ?)`
	tables := make([]models.Videos, 0)
	res = make([]*responseStruct.Video, 0)
	idSlice := make([]int64, 0)
	rows, err := db.Query(sqlstr0, user_id)
	length := 0
	for rows.Next() {
		var temp models.Videos
		err = rows.Scan(&temp.ID, &temp.AuthorID, &temp.PlayUrl, &temp.CoverUrl, &temp.FavoriteCount, &temp.CommentCount, &temp.Title)
		if err != nil {
			zap.L().Info("查询失败", zap.Error(err))
			return nil, err
		}
		idSlice = append(idSlice, temp.AuthorID)
		tables = append(tables, temp)
		length++
	}
	for i := 0; i < length; i++ {
		var r responseStruct.Video
		r.ID = tables[i].ID
		r.Title = tables[i].Title
		r.FavoriteCount = tables[i].FavoriteCount
		r.CommentCount = tables[i].CommentCount
		r.PlayUrl = tables[i].PlayUrl
		r.CoverUrl = tables[i].CoverUrl
		res = append(res, &r)
	}

	sqlstr2 := `select user_id,name,follow_count,follower_count from users where user_id = ?` //查询用户列表
	sqlstr1 := `select id from follows where user_id=? and target_id =?`
	for i := 0; i < length; i++ {
		var au responseStruct.User
		err = db.Get(&au, sqlstr2, idSlice[i])
		if err != nil {
			zap.L().Info("查询失败", zap.Error(err))
			return nil, err
		}

		res[i].Author = au

		_, err = db.Exec(sqlstr1, user_id, idSlice[i]) //关注
		if err != nil {                                //未关注
			res[i].Author.IsFollow = false
		} else {
			res[i].Author.IsFollow = true //关注
		}
		res[i].IsFavorite = true
	}
	return res, nil
}
