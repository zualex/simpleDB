package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/zualex/simpledb/internal/parser"
)

func main() {
	var sql string
	flag.StringVar(&sql, "sql", "", "SQL string")
	flag.Parse()

	fmt.Println("sql:", sql)

	err := parser.Handle(sql)
	if err != nil {
		log.Fatal(err)
	}
}
