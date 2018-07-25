package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/syucream/jackup/src/spanner2mysql"
	"github.com/syucream/spar/src/parser"
)

func main() {
	pathToSql := flag.String("f", "", "path to Spanner DDL")
	flag.Parse()

	var data []byte
	var err error
	if *pathToSql != "" {
		// Try file option
		data, err = ioutil.ReadFile(*pathToSql)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// Try stdin
		data, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
	}

	stmts := parser.Parse(strings.NewReader(string(data)))

	mysqlStmts, err := spanner2mysql.GetMysqlCreateTables(stmts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mysqlStmts)
}
