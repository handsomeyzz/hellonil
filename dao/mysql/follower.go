package mysql

import (
	"go.uber.org/zap"
	"hellonil/responseStruct"
)

func FollowerList(user_id string) (res []*responseStruct.User, err error) {
	//1.判断用户存在与否
	sqlstr := `select id from account where id =?`
	_, err = db.Exec(sqlstr, user_id)
	if err != nil {
		zap.L().Info("用户不存在", zap.Error(err))
		return nil, err
	}
	sqlstr2 := `select user_id from follows where target_id = ?`
	rows, err := db.Query(sqlstr2, user_id)
	if err != nil {
		zap.L().Info("查询失败", zap.Error(err))
		return nil, err
	}
	rsc := make([]*responseStruct.User, 100)
	for rows.Next() {
		var temp responseStruct.User
		err = rows.Scan(&temp)
		if err != nil {
			zap.L().Info("查询失败，有错误", zap.Error(err))
			return nil, err
		}
		rsc = append(rsc, &temp)
	}
	sqlstr3 := `select id from follows where user_id = ? and target_id=?`
	for i, length := 0, len(rsc); i < length; i++ {
		_, err = db.Exec(sqlstr3, user_id, rsc[i].ID)
		if err != nil {
			rsc[i].IsFollow = false
		} else {
			rsc[i].IsFollow = true
		}
	}

	return rsc, nil

}
