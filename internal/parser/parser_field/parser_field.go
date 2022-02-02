package parser_field

import (
	"strings"
)

func GetFieldsFromString(fieldString string) []string {
	fields := strings.Split(strings.TrimSpace(fieldString), ",")

	result := []string{}
	value := ""
	for i := range fields {
		value = strings.TrimSpace(fields[i])
		if value != "" {
			result = append(result, value)
		}
	}

	return result
}
