package ghelper

import (
	"fmt"
	"reflect"
	"strings"
)

func word(text string, out *[]string) {
	pos := strings.LastIndexFunc(text, func(r rune) bool {
		// 'A'-'Z'
		if r >= 65 && r <= 90 {
			return true
		}
		return false
	})
	if pos != -1 {
		*out = append(*out, text[pos:])
		word(text[:pos], out)
	}
}

func sqlName(name string) string {
	var ret string = ""
	var result []string
	word(name, &result)
	for i := len(result) - 1; i >= 0; i-- {
		ret += strings.ToLower(result[i]) + "_"
	}
	return strings.TrimRight(ret, "_")
}

func GInsert(a interface{}) string {
	anyType := reflect.TypeOf(a)
	sql := fmt.Sprintf(`INSERT INTO "%s" (`, sqlName(anyType.Name()))

	var placeholder string
	var index int = 0
	for i := 0; i < anyType.NumField(); i++ {
		filed := anyType.Field(i)
		_, flag := filed.Tag.Lookup("pgsql")
		if flag {
			index++
			sql += fmt.Sprintf(`"%s",`, sqlName(filed.Name))
			placeholder += fmt.Sprintf(`$%d,`, index)
		}
	}

	sql = sql[:len(sql)-1]
	sql += ") VALUES ("

	placeholder = placeholder[:len(placeholder)-1]
	sql += placeholder
	sql += ");"

	return sql
}

// GUpdate
// pgsql:"set"
// pgsql:"where"
func GUpdate(a interface{}) string {
	anyType := reflect.TypeOf(a)

	set := make([]string, 0)
	where := make([]string, 0)
	for i := 0; i < anyType.NumField(); i++ {
		filed := anyType.Field(i)
		tag, flag := filed.Tag.Lookup("pgsql")
		if flag {
			switch tag {
			case "set":
				set = append(set, sqlName(filed.Name))
				break
			case "where":
				where = append(where, sqlName(filed.Name))
				break
			}
		}
	}

	var index int = 0
	sql := fmt.Sprintf(`UPDATE "%s" SET `, sqlName(anyType.Name()))
	for _, s := range set {
		index++
		sql += fmt.Sprintf(`"%s" = $%d,`, s, index)
	}
	if len(set) > 0 {
		sql = sql[:len(sql)-1]
	}

	if len(where) > 0 {
		sql += " WHERE "
	}
	for _, s := range where {
		index++
		sql += fmt.Sprintf(`"%s" = $%d AND `, s, index)
	}
	if len(where) > 0 {
		sql = sql[:len(sql)-5]
	}
	sql += ";"
	return sql
}

// GDelete
// pgsql:"where"
func GDelete(a interface{}) string {
	anyType := reflect.TypeOf(a)

	where := make([]string, 0)
	for i := 0; i < anyType.NumField(); i++ {
		filed := anyType.Field(i)
		tag, flag := filed.Tag.Lookup("pgsql")
		if flag {
			switch tag {
			case "where":
				where = append(where, sqlName(filed.Name))
				break
			}
		}
	}

	var index int = 0
	sql := fmt.Sprintf(`DELETE FROM "%s" `, sqlName(anyType.Name()))
	if len(where) > 0 {
		sql += " WHERE "
	}
	for _, s := range where {
		index++
		sql += fmt.Sprintf(`"%s" = $%d AND `, s, index)
	}
	if len(where) > 0 {
		sql = sql[:len(sql)-5]
	}
	sql += ";"
	return sql
}
