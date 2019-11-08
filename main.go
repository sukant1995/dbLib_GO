package dblib

import (
	"fmt"
	"reflect"
	"strings"
)

// Send is data container for MysqlTable Data which transforms
// slice of any structure to MysqslTable->Data field
//var Send []interface{}

type MysqlTable struct {
	TableName string
	Data      interface{}
}

func prepare(data *MysqlTable) (columns []string, values []string) {
	val := reflect.ValueOf(data.Data)
	for i := 0; i < val.Type().NumField(); i++ {
		columns = append(columns, val.Type().Field(i).Tag.Get("db"))
		v := val.Field(i).Interface()
		str := fmt.Sprintf("%v", v)
		values = append(values, str)
	}
	return columns, values
}

func Insert(data *MysqlTable) error {
	c, v := prepare(data)
	cStr := strings.Join(c, ",")
	vStr := "'" + strings.Join(v, "','") + "'"

	fmt.Println("INSERT INTO " + data.TableName + "(" + cStr + ") VALUES(" + vStr + ")")
	return nil
}

// func getColumns() {}

// func geValue() {

// }
