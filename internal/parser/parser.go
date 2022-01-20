package parser

import (
	"strings"
)

func IsInsert(sql string) bool {
	sqlLowerCase := strings.ToLower(sql)

	indexInsert := strings.Index(sqlLowerCase, "insert")
	if indexInsert == 0 {
		return true
	}

	return false
}
