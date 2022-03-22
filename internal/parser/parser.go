package parser

import (
	"errors"
	"strings"

	"github.com/zualex/simpledb/internal/parser/create_table"
	"github.com/zualex/simpledb/internal/parser/ddl"
	"github.com/zualex/simpledb/internal/parser/insert"
)

func Handle(sql string) error {
	sqlLowerCase := strings.ToLower(sql)

	if ddl.IsCreateTable(sqlLowerCase) {
		return create_table.Handle(sqlLowerCase)
	}

	if ddl.IsInsert(sqlLowerCase) {
		return insert.Handle(sqlLowerCase)
	}

	return errors.New("Некорректный SQL запрос")
}
