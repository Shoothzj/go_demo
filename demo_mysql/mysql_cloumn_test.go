package demo_mysql

import (
	"database/sql"
	"fmt"
	"go_demo/component_mysql/mysql"
	"log"
	"testing"
)

/*
https://golangbot.com/connect-create-db-mysql/
*/
func TestColumn(t *testing.T) {
	db, err := sql.Open("mysql", mysql.Dsn())
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}
	log.Println("Open db success")
	defer db.Close()
	results, err := db.Query("show status")
	if err != nil {
		panic(err)
	}
	defer results.Close()
	log.Println(results)
	for results.Next() {
		columns, err := results.Columns()
		if err != nil {
			log.Println(err)
		}
		for _, column := range columns {
			fmt.Println(column)
		}
		fmt.Println(columns)
	}
}
