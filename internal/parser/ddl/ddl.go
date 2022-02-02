package ddl

import "strings"

const CREATE_TABLE = "create table"
const DROP_TABLE = "drop table"
const INSERT = "insert"
const SELECT = "select"
const UPDATE = "update"
const DELETE = "delete"

func IsCreateTable(sql string) bool {
	return isByDdlCommand(sql, CREATE_TABLE)
}

func isByDdlCommand(sql string, ddlCommand string) bool {
	tokenFields := strings.Fields(ddlCommand)
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
