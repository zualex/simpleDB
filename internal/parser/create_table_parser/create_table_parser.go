package create_table_parser

import (
	"fmt"
	"regexp"
)

func Handle(sqlLowerCase string) {
	// tokens := strings.Fields(sqlLowerCase)
	// fmt.Println(len(tokens))
	// fmt.Println(tokens[0])
	// fmt.Println(tokens[1])
	// fmt.Println(tokens[2])
	// fmt.Println(tokens[3])
	// fmt.Println(tokens[4])
	// fmt.Println(tokens[5])

	isValid(sqlLowerCase)
}

func isValid(sql string) bool {
	// TODO получить группы https://pkg.go.dev/regexp#Regexp.SubexpNames
	match, ok := regexp.MatchString(`^create table(.*?)\((.*?)\);$`, sql)
	fmt.Println(match, ok)

	return false
}
