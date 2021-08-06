package component_hc_config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os/user"
)

type HcConfig struct {
	Region      string `yaml:"region"`
	Ak          string `yaml:"ak"`
	Sk          string `yaml:"sk"`
	ProjectId   string `yaml:"project_id"`
	ObsEndpoint string `yaml:"obs_endpoint"`
	DeviceId    string `yaml:"device_id"`
	DeviceSk    string `yaml:"device_sk"`
}

func ReadHcConfig() *HcConfig {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	homeDirectory := user.HomeDir
	fmt.Printf("Home Directory: %s\n", homeDirectory)
	hcConfig := HcConfig{}
	bytes, err := ioutil.ReadFile(homeDirectory + "/.sh/hc.yaml")
	err = yaml.Unmarshal(bytes, &hcConfig)
	if err != nil {
		panic(err)
	}
	return &hcConfig
}
