package parser

import (
	"strings"
)

const CREATE_TABLE = "create table"
const DROP_TABLE = "drop table"
const INSERT = "insert"
const SELECT = "select"
const UPDATE = "update"
const DELETE = "delete"

func isByToken(sql string, token string) bool {
	tokenFields := strings.Fields(token)
	lenToken := len(tokenFields)

	sqlFields := strings.Fields(strings.ToLower(sql))
	lenSqlField := len(sqlFields)

	if lenToken > lenSqlField {
		return false
	}

	for key := range tokenFields {
		currentToken := tokenFields[key]
		currentSqlField := sqlFields[key]

		if currentToken != currentSqlField {
			return false
		}
	}

	return true
}

func IsCreateTable(sql string) bool {
	return isByToken(sql, CREATE_TABLE)
}

func IsDropTable(sql string) bool {
	return isByToken(sql, DROP_TABLE)
}

func IsInsert(sql string) bool {
	return isByToken(sql, INSERT)
}

func IsSelect(sql string) bool {
	return isByToken(sql, SELECT)
}

func IsUpdate(sql string) bool {
	return isByToken(sql, UPDATE)
}

func IsDelete(sql string) bool {
	return isByToken(sql, DELETE)
}
