package demo_aws

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"testing"
)

func TestStartEc2(t *testing.T) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(hkRegion))
	if err != nil {
		panic(err)
	}
	svc := ec2.NewFromConfig(cfg)
	instancesInput := &ec2.StartInstancesInput{
		InstanceIds:    nil,
		AdditionalInfo: nil,
		DryRun:         false,
	}
	instances, err := svc.StartInstances(context.TODO(), instancesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(instances)
}

func TestCreateEc2Template(t *testing.T) {
	svc := ec2.NewFromConfig(*awsConfig())
	templateName := "kubernetes"
	templateData := &types.RequestLaunchTemplateData{
		BlockDeviceMappings:               nil,
		CapacityReservationSpecification:  nil,
		CpuOptions:                        nil,
		CreditSpecification:               nil,
		DisableApiTermination:             false,
		EbsOptimized:                      false,
		ElasticGpuSpecifications:          nil,
		ElasticInferenceAccelerators:      nil,
		EnclaveOptions:                    nil,
		HibernationOptions:                nil,
		IamInstanceProfile:                nil,
		ImageId:                           nil,
		InstanceInitiatedShutdownBehavior: "",
		InstanceMarketOptions:             nil,
		InstanceType:                      "",
		KernelId:                          nil,
		KeyName:                           nil,
		LicenseSpecifications:             nil,
		MetadataOptions:                   nil,
		Monitoring:                        nil,
		NetworkInterfaces:                 nil,
		Placement:                         nil,
		RamDiskId:                         nil,
		SecurityGroupIds:                  nil,
		SecurityGroups:                    nil,
		TagSpecifications:                 nil,
		UserData:                          nil,
	}
	input := &ec2.CreateLaunchTemplateInput{
		LaunchTemplateData: templateData,
		LaunchTemplateName: &templateName,
		ClientToken:        nil,
		DryRun:             false,
		TagSpecifications:  nil,
		VersionDescription: nil,
	}
	template, err := svc.CreateLaunchTemplate(context.TODO(), input)
	if err != nil {
		panic(err)
	}
	fmt.Println(template)
}

func TestQueryEc2(t *testing.T) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(hkRegion))
	if err != nil {
		panic(err)
	}
	svc := ec2.NewFromConfig(cfg)
	input := &ec2.DescribeInstancesInput{
		DryRun: false,
	}
	describeInstances, err := svc.DescribeInstances(context.TODO(), input)
	if err != nil {
		panic(err)
	}
	fmt.Println(describeInstances)
}
