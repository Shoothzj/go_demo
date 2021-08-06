package mysql

import (
	"database/sql"
	demo_data_gen "demo_data-gen"
	"go_demo/component_mysql/mysql"
	"strconv"
	"testing"
)

func TestFillDateCount(t *testing.T) {
	db, err := sql.Open("mysql", mysql.Dsn())
	if err != nil {
		panic(err)
	}
	sql := "INSERT INTO player_times(player_id, player_time) VALUES (?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	for i := 0; i < 2000; i++ {
		_, err := stmt.Exec("str"+strconv.Itoa(i), demo_data_gen.RandomTimeToday())
		if err != nil {
			panic(err)
		}
	}
}
