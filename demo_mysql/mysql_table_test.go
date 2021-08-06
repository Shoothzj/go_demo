package demo_mysql

import (
	"database/sql"
	"go_demo/component_mysql/mysql"
	"log"
	"testing"
)

func TestCreateTable(t *testing.T) {
	db, err := sql.Open("mysql", mysql.Dsn())
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}
	log.Println("Open db success")
	defer db.Close()
	results, err := db.Exec("CREATE TABLE water_meter(water_meter_id int primary key auto_increment, water_meter_name text, location VARCHAR(30) not null default 'UNKNOWN');")
	if err != nil {
		panic(err)
	}
	rows, err := results.RowsAffected()
	if err != nil {
		panic(err)
	}
	log.Printf("Rows affected when creating table %d\n", rows)
}
