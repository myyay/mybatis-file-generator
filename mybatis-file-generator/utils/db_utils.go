package utils

import (
	"io"
	"log"
	"regexp"
	"strings"
)

type MysqlTable struct {
	TableName string
	Columns   []MysqlColumn
}

type MysqlColumn struct {
	Field   string
	Type    string
	Length  int
	Null    bool
	Key     string
	Default string
	Extra   string
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
