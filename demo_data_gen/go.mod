module demo_data-gen

go 1.16

require (
	github.com/go-sql-driver/mysql v1.6.0
	go_demo/component_mongo v0.0.0
	go_demo/component_mysql v0.0.0
	go_demo/component_util v0.0.0
	go.mongodb.org/mongo-driver v1.5.3
)

replace (
	go_demo/component_mongo v0.0.0 => ../component_mongo
	go_demo/component_mysql v0.0.0 => ../component_mysql
	go_demo/component_util v0.0.0 => ../component_util
)
