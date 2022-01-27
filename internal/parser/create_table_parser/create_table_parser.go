package create_table_parser

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/zualex/simpledb/internal/storage"
)

func Handle(sqlLowerCase string) error {
	isValid, table, fields := getToken(sqlLowerCase)
	if isValid == false {
		return errors.New("Не смог распарсить CREATE TABLE")
	}

	fmt.Println(table)
	fmt.Println(fields)

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
	fieldString := getFields(sqlFields[2])

	return true, tableName, fieldString
}

func getFields(fieldString string) []string {
	fields := strings.Split(strings.TrimSpace(fieldString), ",")

	for i := range fields {
		fields[i] = strings.TrimSpace(fields[i])
	}

	return fields
}
