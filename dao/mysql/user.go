package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"go.uber.org/zap"
	"hellonil/models"
	"hellonil/pkg/snowflake"
	"hellonil/responseStruct"
)

const md5sectret = "hellonil"

// md5加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(md5sectret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// CommentAction

// 判断用户存在与否,存在返回true
func CheckUserExist(username string) bool {
	sqlStr := `select id from accounts where username = ?`
	var id int64
	err := db.Get(&id, sqlStr, username)
	if err == nil {
		//说明用户存在
		return true
	}
	return false
}

// 插入进入数据库
func InsertAccounts(user *models.Accounts) (err error) {
	userID := snowflake.GenID() //雪花算法生成用户id
	user.ID = userID
	pd := encryptPassword(user.Password) //md5加密密码
	sqlStr := `insert into accounts(id,username,password) VALUES (?,?,?)`
	_, err = db.Exec(sqlStr, userID, user.Username, pd)
	if err != nil {
		return err
	}
	return nil
}

// 检查密码
func CheckPassWord(user *models.Accounts) error {
	password := user.Password
	sqlstr := "select id,username,password from accounts where username=?"
	err := db.Get(user, sqlstr, user.Username)
	if err != nil {
		return nil
	}
	if encryptPassword(password) != user.Password {
		return errors.New("密码错误！")
	}
	return nil
}

// 插入user数据进数据库
func InsertUsers(user *models.Accounts) (err error) {
	//生成id
	userid := snowflake.GenID()
	sqlStr := `insert into users(id,user_id,name,avatar,follow_count,follower_count) VALUES (?,?,?,?,?,?)`
	_, err = db.Exec(sqlStr, userid, user.ID, user.Username, "", 0, 0)
	if err != nil {
		return err
	}
	return nil
}

func PublishList(user_id int64) (vlist []*responseStruct.Video, err error) {
	var u responseStruct.User
	sqlstr1 := `select user_id,name,follow_count,follower_count from users where user_id=?`
	err = db.Get(&u, sqlstr1, user_id)
	if err != nil {
		zap.L().Info("查询用户信息失败", zap.Error(err))
		return nil, err
	}

	sqlstr := `select id,play_url,cover_url,favorite_count,comment_count,title from videos where author_id = ?`
	rows, err := db.Query(sqlstr, user_id)
	if err != nil {
		zap.L().Info("发布信息查询失败:", zap.Error(err))
		return nil, err
	}
	length := 0
	vlist = make([]*responseStruct.Video, 0, 100)
	for rows.Next() {
		var temp responseStruct.Video
		err = rows.Scan(&temp.ID, &temp.PlayUrl, &temp.CoverUrl, &temp.FavoriteCount, &temp.CommentCount, &temp.Title)
		temp.Author = u
		if err != nil {
			zap.L().Info("扫描信息失败", zap.Error(err))
			return nil, err
		}
		length++
		vlist = append(vlist, &temp)
	}
	return vlist, nil
}

func SearchUserMsg(uid int) (userMsg *responseStruct.User, err error) {
	sql := `select user_id,name,follow_count,follower_count from users where user_id=?`
	var reu responseStruct.User
	err = db.Get(&reu, sql, uid)
	if err != nil {
		zap.L().Info("用户数据查询失败！错误为：", zap.Error(err))
		return nil, err
	}
	return &reu, nil
}
