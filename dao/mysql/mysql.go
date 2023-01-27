package mysql

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"hellonil/config"
)

var DB *sqlx.DB

func Init() error {
	MysqlX := config.MysqlX()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		MysqlX.User, MysqlX.Password, MysqlX.Addr, MysqlX.Port, MysqlX.DB)
	DB, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return errors.New("mysql初始化失败")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(100)
	return nil
}

func Close() {
	_ = DB.Close()
}
