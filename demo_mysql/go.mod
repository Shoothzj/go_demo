module go_demo/demo_mysql

go 1.17

require (
	github.com/go-sql-driver/mysql v1.6.0
	go_demo/component_mysql v0.0.0
)

replace go_demo/component_mysql v0.0.0 => ../component_mysql