package demo_ucloud

import "github.com/ucloud/ucloud-sdk-go/services/uhost"

func listEcs() []uhost.UHostInstanceSet {
	uHostClient := newHostCli()
	describeUHostInstanceRequest := uHostClient.NewDescribeUHostInstanceRequest()
	describeUHostInstanceRequest.Region = String(region)
	describeUHostInstance, err := uHostClient.DescribeUHostInstance(describeUHostInstanceRequest)
	if err != nil {
		panic(err)
	}
	return describeUHostInstance.UHostSet
}
