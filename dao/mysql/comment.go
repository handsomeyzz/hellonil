package mysql

import (
	"go.uber.org/zap"
	"hellonil/pkg/snowflake"
	"hellonil/responseStruct"
	"time"
)

func comment(username string, vid int) (ru *responseStruct.User, err error) {
	sqlstr2 := `select user_id,name,follow_count,follower_count from users where user_id = (select user_id from accounts where username = ?)`
	var u responseStruct.User
	err = db.Get(&u, sqlstr2, username)
	if err != nil {
		zap.L().Info("评论数据搜索失败，请稍后重试：", zap.Error(err))
		return nil, err
	}
	var isfollow bool
	sqlstr3 := `select target_id from follows where user_id=(select id from accounts where username=?) and target_id =(select author_id from videos where id=?)`
	_, err = db.Exec(sqlstr3, username, vid)
	if err != nil {
		isfollow = false
	} else {
		isfollow = true
	}
	u.IsFollow = isfollow
	return &u, nil
}

func CommentAction(username, ctext string, vid int) (rsc *responseStruct.Comment, err error) {
	id := snowflake.GenID() //生成评论主键
	sqlstr1 := `insert into comments(id,video_id,user_id,content,create_date) VALUES (?,?,(select id from accounts where username=?),?,?)`
	tn := time.Now().Format("2006-01-02 15:04:05")
	_, err = db.Exec(sqlstr1, id, vid, username, ctext, tn)
	if err != nil {
		zap.L().Info("评论数据插入失败，请稍后重试", zap.Error(err))
		return nil, err
	}
	ru, err := comment(username, vid)
	if err != nil {
		return nil, err
	}
	//数据库中videos下的comment_count++
	sqlstr2 := `update videos set comment_count=comment_count+1 where id=?`
	_, err = db.Exec(sqlstr2, vid)
	if err != nil {
		zap.L().Info("更新数据库评论数量失败", zap.Error(err))
		return nil, err
	}

	return &responseStruct.Comment{
		ID:         id,
		User:       *ru,
		Content:    ctext,
		CreateDate: tn,
	}, err
}

func DeleteAction(username string, vid, comment_id int) (rsc *responseStruct.Comment, err error) {
	ru, err := comment(username, vid)
	if err != nil {
		return nil, err
	}
	sqlstr := `delete from comments where id= ?`
	_, err = db.Exec(sqlstr, comment_id)
	if err != nil {
		zap.L().Info("用户信息删除失败，错误信息为：", zap.Error(err))
		return nil, err
	}
	sqlstr2 := `update videos set comment_count=comment_count-1 where id=?`
	_, err = db.Exec(sqlstr2, vid)
	if err != nil {
		zap.L().Info("更新数据库评论数量失败", zap.Error(err))
		return nil, err
	}

	return &responseStruct.Comment{
		ID:         int64(comment_id),
		User:       *ru,
		Content:    "",
		CreateDate: time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}

func CommentIsExist(cid int) bool {
	sqlstr := `select id from comments where id=?`
	_, err := db.Exec(sqlstr, cid)
	if err != nil {
		return false
	}
	return true
}
