package utils

import (
	"io"
	"log"
	"strings"
)

type TableName struct {
	TableName string
}

type MysqlColumn struct {
	Field   string `db:"Field"`
	Type    string `db:"Type"`
	Null    bool   `db:"Null"`
	Key     string `db:"Key"`
	Default string `db:"Default"`
	Extra   string `db:"Extra"`
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
