package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"gopa/dao"
	"log"
)

var id = flag.Int("id", 1, "The Entry ID")

func main() {
	flag.Parse()
	entry, err := dao.GetEntryById(*id)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			fmt.Println(err)
		}
		log.Fatalf("%+v", err)
	}
	fmt.Println("Entry:", entry)
}
