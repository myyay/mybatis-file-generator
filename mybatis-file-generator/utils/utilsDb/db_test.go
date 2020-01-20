package utilsDb

import (
	"fmt"
	"regexp"
	"testing"
)

func TestCompile(t *testing.T) {

	compile := regexp.MustCompile("^(\\w+)\\(?(\\d*,?\\d*)\\)? ?(\\w*)$")

	id := "int(10) unsigned"
	name := "varchar(50)"
	createDate := "datetime"
	updateDate := "timestamp"
	amount := "decimal(22,2)"
	version := "bigint(20)"

	fmt.Println("id", compile.FindStringSubmatch(id))
	fmt.Println("name", compile.FindStringSubmatch(name))
	fmt.Println("createDate", compile.FindStringSubmatch(createDate))
	fmt.Println("updateDate", compile.FindStringSubmatch(updateDate))
	fmt.Println("amount", compile.FindStringSubmatch(amount))
	fmt.Println("version", compile.FindStringSubmatch(version))

	fmt.Println(MysqlTable{})

}
