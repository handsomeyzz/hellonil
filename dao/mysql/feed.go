package mysql

import (
	"go.uber.org/zap"
	"hellonil/models"
	"hellonil/pkg/jwt"
	"hellonil/responseStruct"
	"math/rand"
	"time"
)

func SearchFeed(isOk bool, token string) (res []*responseStruct.Video, err error) {
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
			return nil, err
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
	res = make([]*responseStruct.Video, 0, 100)
	for k, v := range VidAid {
		temp := responseStruct.Video{
			ID: k,
			Author: responseStruct.User{
				ID:            v,
				Name:          u_u[v].Name,
				FollowCount:   u_u[v].FollowCount,
				FollowerCount: u_u[v].FollowerCount,
				IsFollow:      false, //暂时为false
			},
			PlayUrl:       v_v[k].PlayUrl,
			CoverUrl:      v_v[k].CoverUrl,
			FavoriteCount: v_v[k].FavoriteCount,
			CommentCount:  v_v[k].CommentCount,
			IsFavorite:    false, //暂时为false
			Title:         v_v[k].Title,
		}

		res = append(res, &temp)
	}
	//3.当前用户是否有token
	if isOk == false {
		if len(res) >= 30 {
			res = res[:30]
		}
	} else { //有token,解析token
		myc, err := jwt.ParseToken(token)
		if err != nil {
			zap.L().Info("解析token出错！")
			return nil, err
		}
		username := myc.Username
		SearchIsLike(res, username)
		SearchIsFollow(res, username)
		if len(res) >= 30 {
			rand.Seed(time.Now().Unix())
			length := len(res) - 30
			randomMath := rand.Intn(length)
			res = res[randomMath : randomMath+30]
		}
	}

	return res, nil
}

// 查询用户对视频是否喜爱
func SearchIsLike(myfeed []*responseStruct.Video, username string) {
	sqlstr1 := `select video_id from likes where user_id= (SELECT id  FROM accounts WHERE username=?)`
	likeList := make([]int64, 100) //当前用户的喜欢列表
	rows, err := db.Query(sqlstr1, username)
	if err != nil {
		zap.L().Info("查询时发生错误！")
		return
	}
	length := 0
	for rows.Next() {
		err = rows.Scan(&likeList[length])
		if err != nil {
			zap.L().Info("数据扫描时导入发生错误！")
			return
		}
		length++
	}
	for i, t := 0, len(myfeed); i < t; i++ {
		for j := 0; j < length; j++ {
			if likeList[j] == myfeed[i].ID {
				myfeed[i].IsFavorite = true
			}
		}
	}
}

// 查询用户对作者是否关注
func SearchIsFollow(myfeed []*responseStruct.Video, username string) {
	sqlStr := `select target_id from follows where user_id=(SELECT id  FROM accounts WHERE username=?)`
	followsList := make([]int64, 100) //当前用户的喜欢列表
	rows, err := db.Query(sqlStr, username)
	if err != nil {
		zap.L().Info("查询时发生错误！")
		return
	}
	length := 0
	for rows.Next() {
		err = rows.Scan(&followsList[length])
		if err != nil {
			zap.L().Info("数据扫描时导入发生错误！")
			return
		}
		length++
	}
	for i, t := 0, len(myfeed); i < t; i++ {
		for j := 0; j < length; j++ {
			if followsList[j] == myfeed[i].Author.ID {
				myfeed[i].Author.IsFollow = true //is_follow为true
			}
		}
	}
}

// 查询关注数量
func SearchFollowCount(myfeed []*responseStruct.Video, username string) {
	sqlstr := `select count(target_id) from follows where  target_id=(select id from accounts where username =?)`
	var count int64
	rows, err := db.Query(sqlstr, username)
	if err != nil {
		zap.L().Info("查询函数SearchFollowCount查询时发生错误！")
		return
	}
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			zap.L().Info("查询函数SearchFollowCount数据扫描时导入发生错误！")
			return
		}
	}
	for i, t := 0, len(myfeed); i < t; i++ {
		if myfeed[i].Author.Name == username {
			myfeed[i].Author.FollowerCount = count
		}
	}
}
