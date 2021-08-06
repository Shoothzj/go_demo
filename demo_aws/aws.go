package demo_aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

const hkRegion = "ap-east-1"

func awsConfig() *aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(hkRegion))
	if err != nil {
		panic(err)
	}
	return &cfg
}
