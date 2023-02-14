package mysql

import (
	"go.uber.org/zap"
	"hellonil/pkg/snowflake"
	"hellonil/responseStruct"
)

func FollowCheckUid(to_user_id string) bool {
	sqlstr := `select id from accounts where id=?`
	_, err := db.Exec(sqlstr, to_user_id)
	if err != nil {
		return false
	}
	return true
}

func FollowInsert(username string, to_user_id string) (err error) {
	id := snowflake.GenID()
	sqlstr := `insert into follows(id,user_id,target_id) values(?,(select id from accounts where username=?),?)`
	_, err = db.Exec(sqlstr, id, username, to_user_id)
	if err != nil {
		zap.L().Info("关注时插入数据失败", zap.Error(err))
		return err
	}
	//关注数+1
	sqlstr1 := `update users set follow_count = follow_count+1 where user_id=(select id from accounts where username = ?)`
	_, err = db.Exec(sqlstr1, username)
	if err != nil {
		zap.L().Info("用户信息更新失败，请重新操作", zap.Error(err))
		return err
	}

	sqlstr2 := `update users set follower_count = follower_count+1 where user_id=?`
	_, err = db.Exec(sqlstr2, to_user_id)
	if err != nil {
		zap.L().Info("用户信息更新失败，请重新操作", zap.Error(err))
		return err
	}
	return nil
}

func FollowDelete(username string, to_user_id string) (err error) {
	tx := db.MustBegin()
	sqlstr := `delete from follows where user_id = (select user_id from accounts where username =?) and target_id=?`
	sqlstr1 := `update users  set follow_count = follow_count-1 where user_id=(select user_id from accounts where username = ?)`
	sqlstr2 := `update users set follower_count = follower_count-1 where user_id=?`
	tx.MustExec(sqlstr, username, to_user_id)
	tx.MustExec(sqlstr1, username)
	tx.MustExec(sqlstr2, to_user_id)
	err = tx.Commit()
	if err != nil {
		zap.L().Info("关注删除失败")
		return err
	}
	return nil
}

// 关注列表
func FollowList(user_id string) (rsc []*responseStruct.User, err error) {
	sqlstr2 := `select target_id from follows where user_id=?`
	rows, err := db.Query(sqlstr2, user_id)
	if err != nil {
		zap.L().Info("查询失败，请稍后重试", zap.Error(err))
		return nil, err
	}
	res := make([]*responseStruct.User, 0)
	for rows.Next() {
		var temp responseStruct.User
		var tid string
		err = rows.Scan(&tid)
		if err != nil {
			zap.L().Info("信息扫描失败", zap.Error(err))
			return nil, err
		}
		sqlstr := `select user_id,name,follow_count,follower_count from users where user_id = ?`
		err = db.Get(&temp, sqlstr, tid)
		if err != nil {
			zap.L().Info("信息扫描失败", zap.Error(err))
			return nil, err
		}
		temp.IsFollow = true
		res = append(res, &temp)
	}
	return res, nil
}
