package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"myyay/mybatis-file-generator/conf"
	"myyay/mybatis-file-generator/connection"
	"myyay/mybatis-file-generator/utils/utilsDb"
	"myyay/mybatis-file-generator/utils/utilsErr"
	"myyay/mybatis-file-generator/utils/utilsStr"
	"os"
	"strconv"
	"strings"
	"text/template"
)

func main() {

	//defer utils.TryRecover()

	db, err := connection.GetConn(conf.DbConfig)
	utilsErr.LogFatal("connect to db", err)
	defer utilsDb.CloseQuietly(db)
	//获取有哪些表
	tables := getTables(db)
	//将列信息填写进去
	addColumnInfo(tables, db)
	log.Println(tables)

	templateFile := "./mybatis-file-generator/templates/standard_mapper.tmpl"
	outputPath := "d:/test/"
	mapperPath := ""
	daoPath := ""
	domainPath := ""
	mapperPackage := "mapper"
	daoPackage := "com.yay.dao"
	domainPackage := "com.yay.domain"
	fmt.Println(mapperPath, daoPath, domainPath, mapperPackage, daoPackage, domainPackage)

	_ = os.MkdirAll(outputPath+domainPath+strings.ReplaceAll(domainPackage, ".", "/"), 0777)
	_ = os.MkdirAll(outputPath+daoPath+strings.ReplaceAll(daoPackage, ".", "/"), 0777)
	_ = os.MkdirAll(outputPath+mapperPath+strings.ReplaceAll(mapperPackage, ".", "/"), 0777)

	for i := range tables {
		className := utilsStr.ToCamelStr(tables[i].TableName)
		tables[i].DaoClassName = daoPackage + className + "Dao"
		tables[i].DomainClassName = domainPackage + className

		mapperFileResultPath := outputPath + mapperPath + strings.ReplaceAll(mapperPackage, ".", "/") + "/" + className + "Dao.xml"

		t, err := template.ParseFiles(templateFile)
		utilsErr.LogFatal("parse template failed : "+templateFile, err)
		resultFile, err := os.Create(mapperFileResultPath)
		defer utilsDb.CloseQuietly(resultFile)
		writer := bufio.NewWriter(resultFile)
		utilsErr.LogFatal("create file failed : "+mapperFileResultPath, err)
		err = t.Execute(writer, tables[i])
		utilsErr.LogFatal("write file failed : "+mapperFileResultPath, err)
		writer.Flush()

	}

}

func addColumnInfo(tables []utilsDb.MysqlTable, db *sql.DB) {
	for i := range tables {
		columns, err := db.Query("desc " + tables[i].TableName)
		utilsErr.LogPanic("read Fields", err)

		for columns.Next() {

			var cols utilsDb.MysqlColumnString
			_ = columns.Scan(&cols.Field, &cols.Type, &cols.Null, &cols.Key, &cols.Default, &cols.Extra)
			compileResult, _ := utilsDb.CompileType(cols.Type)
			length := 0
			if len(compileResult) > 1 {
				if strings.Contains(compileResult[1], ",") {
					length, _ = strconv.Atoi(strings.Split(compileResult[1], ",")[0])

				} else {
					length, _ = strconv.Atoi(compileResult[1])
				}
			}

			tables[i].Columns = append(tables[i].Columns, utilsDb.MysqlColumn{
				ColumnName:      cols.Field,
				Type:            compileResult[0],
				Length:          length,
				Null:            strings.EqualFold(cols.Null, "YES"),
				Key:             cols.Key,
				Default:         cols.Default,
				Extra:           cols.Extra,
				JdbcType:        utilsDb.GetJdbcType(compileResult[0]),
				PropertyName:    utilsStr.ToCamelStr2(cols.Field),
				IsPrimary:       strings.EqualFold(cols.Key, "PRI"),
				IsAutoIncrement: strings.EqualFold(cols.Extra, "auto_increment"),
				IsVersion:       strings.EqualFold(cols.Field, "version"),
				IsAutoUpdate:    strings.HasPrefix(cols.Extra, "on update"),
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

func getTables(db *sql.DB) []utilsDb.MysqlTable {
	tablesRes, err := db.Query("show tables")
	utilsErr.LogPanic("show tables", err)

	defer utilsDb.CloseQuietly(tablesRes)
	var tables []utilsDb.MysqlTable
	for tablesRes.Next() {
		var table utilsDb.MysqlTable
		err := tablesRes.Scan(&table.TableName)
		utilsErr.LogPanic("read tables", err)
		tables = append(tables, table)

	}

	return tables
}
