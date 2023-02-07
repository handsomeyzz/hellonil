package mysql

import "go.uber.org/zap"

func UnIsOkUID(username string, user_id int64) bool {
	sqlstr := `select id from accounts where username=? `
	var uid int64
	err := db.Get(&uid, sqlstr, username)
	if err != nil {
		zap.L().Info("当前用户的用户名和用户ID不匹配！", zap.Error(err))
		return false
	}
	return true
}
