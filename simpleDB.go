package main

import (
	"flag"
	"fmt"
)

func main() {
	sqlString := flag.String("sql", "", "SQL string")
	flag.Parse()

	fmt.Println("sql:", *sqlString)
}
