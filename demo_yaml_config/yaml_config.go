package main

import (
	"flag"
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

var yamlConfigFile = flag.String("configFile", "config/example.yaml", "Input your config path")

type T struct {
	Name     string
	Job      string
	Skill    string
	Employed string
}

func main() {
	flag.Parse()
	fmt.Println(*yamlConfigFile)
	bytes, err := ioutil.ReadFile(*yamlConfigFile)
	check(err)
	t := T{}
	err = yaml.Unmarshal(bytes, &t)
	check(err)
	fmt.Println(t)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
