package mysql

import (
	"database/sql"
	"demo_data-gen"
	"go_demo/component_mysql/mysql"
	"strconv"
	"testing"
)

func TestFillGeoData(t *testing.T) {
	db, err := sql.Open("mysql", mysql.Dsn())
	if err != nil {
		panic(err)
	}
	sql := "INSERT INTO player_geo(player_id, longtitude, latitude) VALUES (?, ?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	for i := 0; i < 2000; i++ {
		_, err := stmt.Exec("str"+strconv.Itoa(i), demo_data_gen.RandomFloat64FromRange(73.40, 135.2), demo_data_gen.RandomFloat64FromRange(3, 53))
		if err != nil {
			panic(err)
		}
	}
}
