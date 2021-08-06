package mysql

import (
	"database/sql"
	_ "database/sql"
	"demo_data-gen"
	"demo_data-gen/comic"
	"demo_data-gen/iso3166"
	_ "github.com/go-sql-driver/mysql"
	"go_demo/component_mysql/mysql"
	"strconv"
	"testing"
)

func TestFillData(t *testing.T) {
	db, err := sql.Open("mysql", mysql.Dsn())
	if err != nil {
		panic(err)
	}
	sql := "INSERT INTO player_now(player_id, game, country, city) VALUES (?, ?, ?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	for i := 0; i < 2000; i++ {
		_, err := stmt.Exec("str"+strconv.Itoa(i), demo_data_gen.RandomString(comic.Comic), demo_data_gen.RandomStringIndex(0, 10, iso3166.CountryCode), demo_data_gen.RandomString(iso3166.ChinaCode))
		if err != nil {
			panic(err)
		}
	}
	for i := 0; i < 2000; i++ {
		_, err := stmt.Exec("str"+strconv.Itoa(i), demo_data_gen.RandomString(comic.Comic), demo_data_gen.RandomStringIndex(0, 50, iso3166.CountryCode), demo_data_gen.RandomString(iso3166.ChinaCode))
		if err != nil {
			panic(err)
		}
	}
	for i := 0; i < 2000; i++ {
		_, err := stmt.Exec("str"+strconv.Itoa(i), demo_data_gen.RandomString(comic.Comic), demo_data_gen.RandomStringIndex(0, 100, iso3166.CountryCode), demo_data_gen.RandomString(iso3166.ChinaCode))
		if err != nil {
			panic(err)
		}
	}
	for i := 0; i < 2000; i++ {
		_, err := stmt.Exec("str"+strconv.Itoa(i), demo_data_gen.RandomString(comic.Comic), demo_data_gen.RandomString(iso3166.CountryCode), demo_data_gen.RandomString(iso3166.ChinaCode))
		if err != nil {
			panic(err)
		}
	}
}
