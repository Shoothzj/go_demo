package demo_mysql

import (
	"database/sql"
	"go_demo/component_mysql/mysql"
	"go_demo/demo_mysql/module"
	"log"
	"testing"
)

/*
https://golangbot.com/connect-create-db-mysql/
*/
func TestDbName(t *testing.T) {
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
		var statusDto module.StatusTable
		// for each row, scan the result into our statusDto composite object
		err = results.Scan(&statusDto.VariableName, &statusDto.Value)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the statusDto's Name attribute
		log.Printf("status key is %s value is %s", statusDto.VariableName, statusDto.Value)
	}
}
