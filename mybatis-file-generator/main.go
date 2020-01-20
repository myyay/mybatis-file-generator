package main

import (
	"database/sql"
	"log"
	"myyay/mybatis-file-generator/conf"
	"myyay/mybatis-file-generator/connection"
	"myyay/mybatis-file-generator/utils"
	"strconv"
	"strings"
)

func main() {

	//defer utils.TryRecover()

	db, err := connection.GetConn(conf.DbConfig)
	utils.LogFatal("connect to db", err)
	defer utils.CloseQuietly(db)
	//获取有哪些表
	tables := getTables(db)
	//将列信息填写进去
	addColumnInfo(tables, db)
	log.Println(tables)

}

func addColumnInfo(tables []utils.MysqlTable, db *sql.DB) {
	for i := range tables {
		columns, err := db.Query("desc " + tables[i].TableName)
		utils.LogPanic("read Fields", err)

		for columns.Next() {

			var cols utils.MysqlColumnString
			_ = columns.Scan(&cols.Field, &cols.Type, &cols.Null, &cols.Key, &cols.Default, &cols.Extra)
			compileResult, _ := utils.CompileType(cols.Type)
			length := 0
			if len(compileResult) > 1 {
				if strings.Contains(compileResult[1], ",") {
					length, _ = strconv.Atoi(strings.Split(compileResult[1], ",")[0])

				} else {
					length, _ = strconv.Atoi(compileResult[1])
				}
			}

			null := false
			if strings.EqualFold(cols.Null, "YES") {
				null = true
			}

			tables[i].Columns = append(tables[i].Columns, utils.MysqlColumn{
				Field:   cols.Field,
				Type:    compileResult[0],
				Length:  length,
				Null:    null,
				Key:     cols.Key,
				Default: cols.Default,
				Extra:   cols.Extra,
			})
		}

	}
}

func makeCacheList(size int) []interface{} {
	cache := make([]interface{}, size)
	for i, _ := range cache {
		var o interface{}
		cache[i] = &o
	}
	return cache
}

func getTables(db *sql.DB) []utils.MysqlTable {
	tablesRes, err := db.Query("show tables")
	utils.LogPanic("show tables", err)

	defer utils.CloseQuietly(tablesRes)
	var tables []utils.MysqlTable
	for tablesRes.Next() {
		var table utils.MysqlTable
		err := tablesRes.Scan(&table.TableName)
		utils.LogPanic("read tables", err)
		tables = append(tables, table)

	}

	return tables
}
