package main

import (
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"testing"
)

const (
	Username     = "hzj"
	Password     = "Mysql@123"
	Hostname     = "127.0.0.1:3306"
	DatabaseName = "ttbb"
)

func TestGetGlobalStatus(t *testing.T) {
	gdb.SetConfig(gdb.Config{
		"default": gdb.ConfigGroup{
			gdb.ConfigNode{
				Host: "localhost",
				Port: "3306",
				User: Username,
				Pass: Password,
				Type: "mysql",
			},
		},
	})
	result, err := g.DB().GetAll("SHOW GLOBAL STATUS")
	if err != nil {
		panic(err)
	}
	resMap := make(map[string]string)
	for _, record := range result {
		m := record.Map()
		fmt.Println(record)
		key, ok := m["Variable_name"].(string)
		if !ok {
			continue
		}
		value, ok := m["Value"].(string)
		if !ok {
			continue
		}
		resMap[key] = value
		fmt.Println(m)
	}
	fmt.Println(resMap)
}
