package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/zualex/simpledb/internal/parser"
	"github.com/zualex/simpledb/internal/parser/create_table_parser"
)

func main() {
	var sql string
	flag.StringVar(&sql, "sql", "", "SQL string")
	flag.Parse()

	fmt.Println("sql:", sql)
	sqlLowerCase := strings.ToLower(sql)

	if parser.IsCreateTable(sql) {
		err := create_table_parser.Handle(sqlLowerCase)
		if err != nil {
			log.Fatal(err)
		}
	}
}
