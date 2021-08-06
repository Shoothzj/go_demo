package mysql

import "fmt"

const (
	Username     = "hzj"
	Password     = "Mysql@123"
	Hostname     = "127.0.0.1:3306"
	DatabaseName = "ttbb"
)

func Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", Username, Password, Hostname, DatabaseName)
}
