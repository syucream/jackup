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
	"github.com/syucream/jackup/src/converter"
)

var initializers = map[string]func(config converter.Config) converter.Converter{
	"spanner2mysql": spanner2mysql.NewSpanner2MysqlConverter,
	// TODO
}

func main() {
	pathToSql := flag.String("f", "", "path to DDL .sql file")
	target := flag.String("t", "spanner2mysql", "convert target (spanner2mysql, ...)")
	strictFlag := flag.Bool("strict", false, "Strict check")
	allowConvertStringFlag := flag.Bool("allow-convert-string", true, "Convert between long string")
	removeIndexNameFlag := flag.Bool("remove-index-name", true, "Remove long index name")

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

	var c converter.Converter
	if initFunc, ok := initializers[*target]; ok {
		c = initFunc(converter.Config{
			Strict:             *strictFlag,
			AllowConvertString: *allowConvertStringFlag,
			RemoveIndexName:    *removeIndexNameFlag,
		})
	} else {
		log.Fatal("Cannot find target converter : %s", *target)
	}

	mysqlStmts, err := c.Convert(stmts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mysqlStmts)
}
