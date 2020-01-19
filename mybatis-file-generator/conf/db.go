package conf

import (
	"fmt"
	"net/url"
)

type DbConf struct {
	Host     string
	Port     int
	UserName string
	Password string
	DataBase string
}

const (
	MysqlDriverName = "mysql"
	Loc             = "Asia/Shanghai"
	mysqlConnStr    = "%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=%s&parseTime=true"
)

var DbConfig DbConf = DbConf{
	Host:     "127.0.0.1",
	Port:     3306,
	UserName: "root",
	Password: "123456",
	DataBase: "yay",
}

func GetMysqlConnStr(dbConf DbConf) string {
	return fmt.Sprintf(mysqlConnStr, dbConf.UserName, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DataBase, url.QueryEscape(Loc))
}
