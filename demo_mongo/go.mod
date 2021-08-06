module go_demo/demo_mongo

go 1.16

require (
	go_demo/component_mongo v0.0.0
	go.mongodb.org/mongo-driver v1.5.3
)

replace go_demo/component_mongo v0.0.0 => ../component_mongo
