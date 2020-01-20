package utilsDb

import (
	"io"
	"log"
	"regexp"
	"strings"
)

type MysqlTable struct {
	TableName          string
	DomainClassPackage string
	DomainClassName    string
	DaoClassNPackage   string
	DaoClassName       string
	Columns            []MysqlColumn
}

type MysqlColumn struct {
	ColumnName      string
	Type            string
	Length          int
	Null            bool
	Key             string
	Default         string
	Extra           string
	JdbcType        string
	PropertyName    string
	IsPrimary       bool
	IsAutoIncrement bool
	IsVersion       bool
	IsAutoUpdate    bool
}

type MysqlColumnString struct {
	Field   string `db:"Field"`
	Type    string `db:"Type"`
	Null    string `db:"Null"`
	Key     string `db:"Key"`
	Default string `db:"Default"`
	Extra   string `db:"Extra"`
}

//Type正则解析
var compile *regexp.Regexp

func init() {
	compile = regexp.MustCompile("^(\\w+)\\(?(\\d*,?\\d*)\\)? ?(\\w*)$")
}

func CompileType(tp string) (str []string, err error) {
	defer func() {
		r := recover()
		if e, ok := r.(error); ok {
			str, err = []string{}, e
		}
	}()

	return compile.FindStringSubmatch(tp)[1:], nil
}

func (c *MysqlColumn) GetTypeName() string {
	return strings.Split(c.Type, "(")[0]

}

func CloseQuietly(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		log.Println("db connection close error", err)
	}
}

var typeForMysqlToJdbcType = map[string]string{
	"int":                "INTEGER",
	"integer":            "INTEGER",
	"tinyint":            "TINYINT",
	"smallint":           "SMALLINT",
	"mediumint":          "INTEGER",
	"bigint":             "BIGINT",
	"int unsigned":       "INTEGER",
	"integer unsigned":   "INTEGER",
	"tinyint unsigned":   "TINYINT",
	"smallint unsigned":  "SMALLINT",
	"mediumint unsigned": "INTEGER",
	"bigint unsigned":    "BIGINT",
	"bit":                "INTEGER",
	"bool":               "Boolean",
	"enum":               "VARCHAR",
	"set":                "VARCHAR",
	"varchar":            "VARCHAR",
	"char":               "CHAR",
	"tinytext":           "VARCHAR",
	"mediumtext":         "VARCHAR",
	"text":               "VARCHAR",
	"longtext":           "VARCHAR",
	"blob":               "VARCHAR",
	"tinyblob":           "VARCHAR",
	"mediumblob":         "VARCHAR",
	"longblob":           "VARCHAR",
	"date":               "DATE",      // time.Time or string
	"datetime":           "DATETIME",  // time.Time or string
	"timestamp":          "TIMESTAMP", // time.Time or string
	"time":               "TIMESTAMP", // time.Time or string
	"float":              "DECIMAL",
	"double":             "DECIMAL",
	"decimal":            "DECIMAL",
	"binary":             "VARCHAR",
	"varbinary":          "VARCHAR",
}

var typeForMysqlToJavaType = map[string]string{
	"int":                "Integer",
	"integer":            "Integer",
	"tinyint":            "Integer",
	"smallint":           "Integer",
	"mediumint":          "Integer",
	"bigint":             "Long",
	"int unsigned":       "Integer",
	"integer unsigned":   "Integer",
	"tinyint unsigned":   "Integer",
	"smallint unsigned":  "Integer",
	"mediumint unsigned": "Integer",
	"bigint unsigned":    "Long",
	"bit":                "Integer",
	"bool":               "Boolean",
	"enum":               "String",
	"set":                "String",
	"varchar":            "String",
	"char":               "String",
	"tinytext":           "String",
	"mediumtext":         "String",
	"text":               "String",
	"longtext":           "String",
	"blob":               "String",
	"tinyblob":           "String",
	"mediumblob":         "String",
	"longblob":           "String",
	"date":               "Date", // time.Time or string
	"datetime":           "Date", // time.Time or string
	"timestamp":          "Date", // time.Time or string
	"time":               "Date", // time.Time or string
	"float":              "BigDecimal",
	"double":             "BigDecimal",
	"decimal":            "BigDecimal",
	"binary":             "String",
	"varbinary":          "String",
}

func GetJdbcType(tp string) string {
	return typeForMysqlToJdbcType[tp]
}
func GetJavaType(tp string) string {
	return typeForMysqlToJdbcType[tp]
}
