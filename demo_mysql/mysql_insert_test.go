package demo_mysql

import (
	"database/sql"
	"go_demo/component_mysql/mysql"
	"go_demo/demo_mysql/module"
	"log"
	"testing"
)

func TestInsertData(t *testing.T) {
	db, err := sql.Open("mysql", mysql.Dsn())
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}
	waterMeter := module.WaterMeter{WaterMeterId: 1, WaterMeterName: "cn"}
	Insert(db, &waterMeter)
}

func Insert(db *sql.DB, meter *module.WaterMeter) {
	sql := "INSERT INTO water_meter(water_meter_id, water_meter_name) VALUES (?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(meter.WaterMeterId, meter.WaterMeterName)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	log.Printf("%d water meter created", rowsAffected)
}
