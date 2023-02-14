package mysql

import (
	"go.uber.org/zap"
	"hellonil/pkg/snowflake"
	"time"
)

func InsertMsg(from_user_id any, to_user_id, content string) (err error) {
	sqlstr := `insert into message(id,from_user_id,to_user_id,content,create_data) values(?,?,?,?,?)`
	id := snowflake.GenID()
	create_data := time.Now().Format("2006-01-02 15:04:05")
	_, err = db.Exec(sqlstr, id, from_user_id, to_user_id, content, create_data)
	if err != nil {
		zap.L().Info("插入发送的消息失败")
		return err
	}
	return nil
}
