package parser

import (
	"strings"
)

const INSERT = "insert"
const SELECT = "select"
const UPDATE = "update"

func IsInsert(sql string) bool {
	sqlLowerCase := strings.ToLower(sql)

	indexInsert := strings.Index(sqlLowerCase, INSERT)
	if indexInsert == 0 {
		return true
	}

	return false
}
