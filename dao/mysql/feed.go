package mysql

import (
	"go.uber.org/zap"
	"hellonil/models"
)

type Video struct {
	ID            int64  `json:"id"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	Title         string `json:"title"`
	IsFavorite    bool   `json:"is_favorite"`
	Author        User   `json:"author"`
}
type User struct {
	ID            int64  `json:"id"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	Name          string `json:"name"`
	IsFollow      bool   `json:"is_follow"`
}

func SearchFeed(myfeed []*Video, isOk bool) (err error) {
	//1.取视频id和作者id
	sqlstr1 := `select id,author_id from videos`
	rows, err := db.Query(sqlstr1)
	VidAid := make(map[int64]int64)
	for rows.Next() {
		var vid int64
		var aid int64
		err = rows.Scan(&vid, &aid)
		VidAid[vid] = aid
	}
	//2.查询作者信息
	var user []models.Users
	u_u := make(map[int64]models.Users, 100)
	sqlstr2 := `select id,user_id,name,avatar,follow_count,follower_count from users where user_id=?`
	for _, v := range VidAid {
		var u []models.Users
		err = db.Select(&u, sqlstr2, v)
		if err != nil {
			zap.L().Info("Feed流视频查询失败1")
			return
		}
		u_u[v] = u[0]
		user = append(user, u[0])
	}
	//3。查询视频信息
	var videos []models.Videos
	v_v := make(map[int64]models.Videos, 100)
	sqlstr3 := `select id,author_id,play_url,cover_url,favorite_count,comment_count,title,create_time from videos where id=?`
	for k, _ := range VidAid {
		var v []models.Videos
		err = db.Select(&v, sqlstr3, k)
		if err != nil {
			zap.L().Info("Feed流视频查询失败2")
			return
		}
		v_v[k] = v[0]
		videos = append(videos, v[0])
	}
	//4.组合
	for k, v := range VidAid {
		temp := Video{
			ID: k,
			Author: User{
				ID:          v,
				Name:        u_u[v].Name,
				FollowCount: u_u[v].FollowCount,
				IsFollow:    false, //暂时为false
			},
			PlayUrl:       v_v[k].PlayUrl,
			CoverUrl:      v_v[k].CoverUrl,
			FavoriteCount: v_v[k].FavoriteCount,
			CommentCount:  v_v[k].CommentCount,
			IsFavorite:    false, //暂时为false
			Title:         v_v[k].Title,
		}
		myfeed = append(myfeed, &temp)
	}
	//3.当前用户是否有token
	if isOk == false {
		if len(myfeed) >= 30 {
			myfeed = myfeed[:30]
		}
	} else { //有token

	}

	return nil
}

func SearchIsLike(myfeed []*Video, username string) {
	//1.查询用户对视频是否喜爱
	//sqlstr1 := `select id from accounts where username = username`

}
