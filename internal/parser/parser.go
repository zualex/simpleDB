package parser

import (
	"errors"
	"strings"

	"github.com/zualex/simpledb/internal/parser/create_table_parser"
)

const CREATE_TABLE = "create table"
const DROP_TABLE = "drop table"
const INSERT = "insert"
const SELECT = "select"
const UPDATE = "update"
const DELETE = "delete"

func Handle(sql string) error {
	sqlLowerCase := strings.ToLower(sql)

	if isByToken(sqlLowerCase, CREATE_TABLE) {
		return create_table_parser.Handle(sqlLowerCase)
	}

	return errors.New("Некорректный SQL запрос")
}

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
