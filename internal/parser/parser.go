package parser

import (
	"errors"
	"strings"

	"github.com/zualex/simpledb/internal/parser/create_table"
	"github.com/zualex/simpledb/internal/parser/ddl"
)

func Handle(sql string) error {
	sqlLowerCase := strings.ToLower(sql)

	if ddl.IsCreateTable(sqlLowerCase) {
		return create_table.Handle(sqlLowerCase)
	}

	return errors.New("Некорректный SQL запрос")
}
