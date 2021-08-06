package demo_ucloud

import (
	"fmt"
	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/services/ulb"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os/user"
	"time"
)

type UcloudConfig struct {
	Ak string `yaml:"ak"`
	Sk string `yaml:"sk"`
}

func Bool(i bool) *bool {
	return &i
}

func Int(i int) *int {
	return &i
}

func Int64(i int64) *int64 {
	return &i
}

func String(i string) *string {
	return &i
}

func Duration(i time.Duration) *time.Duration {
	return &i
}

func readAkSk() *UcloudConfig {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	homeDirectory := user.HomeDir
	fmt.Printf("Home Directory: %s\n", homeDirectory)
	hcConfig := UcloudConfig{}
	bytes, err := ioutil.ReadFile(homeDirectory + "/.sh/ucloud.yaml")
	err = yaml.Unmarshal(bytes, &hcConfig)
	if err != nil {
		panic(err)
	}
	return &hcConfig
}

func newHostCli() *uhost.UHostClient {
	config := ucloud.NewConfig()
	config.Region = region

	// replace the public/private key by your own
	credential := auth.NewCredential()
	akSk := readAkSk()
	credential.PublicKey = akSk.Ak
	credential.PrivateKey = akSk.Sk

	return uhost.NewClient(&config, &credential)
}

func newElbCli() *ulb.ULBClient {
	config := ucloud.NewConfig()
	config.Region = region

	// replace the public/private key by your own
	credential := auth.NewCredential()
	akSk := readAkSk()
	credential.PublicKey = akSk.Ak
	credential.PrivateKey = akSk.Sk

	return ulb.NewClient(&config, &credential)
}