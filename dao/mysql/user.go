package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"hellonil/models"
	"hellonil/pkg/snowflake"
)

const md5sectret = "hellonil"

// md5加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(md5sectret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// 判断用户存在与否,存在返回true
func CheckUserExist(user *models.Accounts) bool {
	sqlStr := `select id,username,password from accounts where username = ?`
	err := db.Get(user, sqlStr, user.Username)
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
