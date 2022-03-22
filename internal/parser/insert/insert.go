package insert

import (
	"fmt"
	"regexp"
	"strings"

	parser_field "github.com/zualex/simpledb/internal/parser/field"
	"github.com/zualex/simpledb/internal/storage"
)

func Handle(sql string) error {
	table, valueMap, error := getToken(sql)
	if error != nil {
		return error
	}

	return storage.Insert(table, valueMap)
}

func getToken(sql string) (string, map[string]string, error) {
	re := regexp.MustCompile(`^insert into (?P<table>.*?)\((?P<fields>.*?)\) values \((?P<values>.*?)\);$`)
	isValid := re.MatchString(sql)
	sqlFields := re.FindStringSubmatch(sql)

	if isValid == false || len(sqlFields) != 4 {
		return "", nil, nil
	}

	tableName := strings.TrimSpace(sqlFields[1])
	fields := parser_field.GetFieldsFromString(sqlFields[2])
	values := parser_field.GetFieldsFromString(sqlFields[3])

	if len(fields) != len(values) {
		return "", nil, fmt.Errorf("Не совпадает кол-во полей (fields: %d) и кол-во значений (values: %d)", len(fields), len(values))
	}

	valueMap := map[string]string{}
	for i, val := range fields {
		valueMap[val] = values[i]
	}

	return tableName, valueMap, nil
}
