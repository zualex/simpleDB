package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/zualex/simpledb/internal/parser"
	// "internal/parser"
)

func ParseInsert(sql string) {
	fmt.Println("INSERT")
}

func main() {
	var sql string
	flag.StringVar(&sql, "sql", "", "SQL string")

	flag.Parse()

	fmt.Println("sql:", sql)
	sqlLowerCase := strings.ToLower(sql)

	if parser.IsInsert(sql) {
		ParseInsert(sqlLowerCase)
	}
}
