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

	stmts, err := parser.Parse(strings.NewReader(string(data)))
	if err != nil {
		log.Fatal(err)
	}

	converter := spanner2mysql.Spanner2MysqlConverter{}
	mysqlStmts, err := converter.Convert(stmts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mysqlStmts)
}
