package demo_hc

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	component_hc_config "go_demo/component_hc_config"
	"strings"
	"testing"
)

func TestCreateBucket(t *testing.T) {
	hcConfig := component_hc_config.ReadHcConfig()
	obsClient, err := obs.New(hcConfig.Ak, hcConfig.Sk, hcConfig.ObsEndpoint)
	if err != nil {
		panic(err)
	}
	input := &obs.CreateBucketInput{}
	input.Bucket = "zj-bucket"
	input.Location = "cn-north-4"
	output, err := obsClient.CreateBucket(input)
	if err == nil {
		fmt.Printf("RequestId:%s\n", output.RequestId)
	} else if obsError, ok := err.(obs.ObsError); ok {
		fmt.Println(obsError.Code)
		fmt.Println(obsError.Message)
	}
}

func TestUploadDataFile(t *testing.T) {
	hcConfig := component_hc_config.ReadHcConfig()
	obsClient, err := obs.New(hcConfig.Ak, hcConfig.Sk, hcConfig.ObsEndpoint)
	if err != nil {
		panic(err)
	}
	input := &obs.PutFileInput{}
	input.Bucket = "zj-bucket"
	input.Key = "zj-object"
	input.SourceFile = "localfile.csv"
	output, err := obsClient.PutFile(input)
	if err == nil {
		fmt.Printf("RequestId:%s\n", output.RequestId)
	} else if obsError, ok := err.(obs.ObsError); ok {
		fmt.Println(obsError.Code)
		fmt.Println(obsError.Message)
	}
}

func TestUploadDataStream(t *testing.T) {
	hcConfig := component_hc_config.ReadHcConfig()
	obsClient, err := obs.New(hcConfig.Ak, hcConfig.Sk, hcConfig.ObsEndpoint)
	if err != nil {
		panic(err)
	}
	input := &obs.PutObjectInput{}
	input.Bucket = "zj-bucket"
	input.Key = "zj-object"
	input.Body = strings.NewReader("Hello OBS")
	output, err := obsClient.PutObject(input)
	if err == nil {
		fmt.Printf("RequestId:%s\n", output.RequestId)
	} else if obsError, ok := err.(obs.ObsError); ok {
		fmt.Println(obsError.Code)
		fmt.Println(obsError.Message)
	}
}

func TestDeleteBucket(t *testing.T) {
	hcConfig := component_hc_config.ReadHcConfig()
	obsClient, err := obs.New(hcConfig.Ak, hcConfig.Sk, hcConfig.ObsEndpoint)
	if err != nil {
		panic(err)
	}
	output, err := obsClient.DeleteBucket("zj-bucket")
	if err == nil {
		fmt.Printf("RequestId:%s\n", output.RequestId)
	} else if obsError, ok := err.(obs.ObsError); ok {
		fmt.Println(obsError.Code)
		fmt.Println(obsError.Message)
	}
}
