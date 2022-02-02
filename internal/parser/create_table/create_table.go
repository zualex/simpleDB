package create_table

import (
	"errors"
	"regexp"
	"strings"

	parser_field "github.com/zualex/simpledb/internal/parser/field"
	"github.com/zualex/simpledb/internal/storage"
)

func Handle(sqlLowerCase string) error {
	isValid, table, fields := getToken(sqlLowerCase)
	if isValid == false {
		return errors.New("Не смог распарсить CREATE TABLE")
	}

	return storage.Create(table, fields)
}

func getToken(sql string) (bool, string, []string) {
	re := regexp.MustCompile(`^create table(?P<table>.*?)\((?P<fields>.*?)\);$`)
	isValid := re.MatchString(sql)
	sqlFields := re.FindStringSubmatch(sql)

	if isValid == false || len(sqlFields) != 3 {
		return false, "", []string{}
	}

	tableName := strings.TrimSpace(sqlFields[1])
	fieldString := parser_field.GetFieldsFromString(sqlFields[2])

	return true, tableName, fieldString
}
