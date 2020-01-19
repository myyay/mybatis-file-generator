package main

import (
	"log"
	"myyay/mybatis-file-generator/conf"
	"myyay/mybatis-file-generator/connection"
	"myyay/mybatis-file-generator/utils"
)

func main() {

	//defer utils.TryRecover()

	db, err := connection.GetConn(conf.DbConfig)
	utils.LogFatal("connect to db", err)
	defer db.Close()
	defer utils.CloseQuietly(db)

	tablesRes, err := db.Query("show tables")
	utils.LogPanic("show tables", err)

	defer utils.CloseQuietly(tablesRes)

	var tables []utils.TableName
	for tablesRes.Next() {
		var tableName utils.TableName
		err := tablesRes.Scan(&tableName.TableName)
		utils.LogPanic("read tables", err)
		tables = append(tables, tableName)

	}

	log.Println("tables", tables)

}
