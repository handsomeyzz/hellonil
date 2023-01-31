package mysql

import (
	"crypto/md5"
	"encoding/hex"
)

const md5sectret = "hellonil"

// md5加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(md5sectret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// 判断用户存在与否
func CheckUserExist(username string) (err error) {
	return nil
}

func InsertUser() (err error) {
	return nil
}
