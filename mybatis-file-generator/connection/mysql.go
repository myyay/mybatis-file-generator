package connection

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"myyay/mybatis-file-generator/conf"
)

func GetConn(dbConf conf.DbConf) (*sql.DB, error) {
	return sql.Open(conf.MysqlDriverName, conf.GetMysqlConnStr(dbConf))
}
