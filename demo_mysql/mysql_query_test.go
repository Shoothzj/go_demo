package demo_mysql

import (
	"database/sql"
	"fmt"
	"go_demo/component_mysql/mysql"
	"go_demo/demo_mysql/module"
	"log"
	"testing"
)

func TestQueryData(t *testing.T) {
	db, err := sql.Open("mysql", mysql.Dsn())
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}
	waterMeter := module.WaterMeter{}
	Query(db, &waterMeter)
}

func Query(db *sql.DB, meter *module.WaterMeter) {
	sql := "SELECT * FROM water_meter WHERE water_meter_id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	query, err := stmt.Query(1)
	for query.Next() {
		fmt.Println("begin to scan")
		err := query.Scan(&meter.WaterMeterId, &meter.WaterMeterName, &meter.Location)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("meter is ", meter)
}
