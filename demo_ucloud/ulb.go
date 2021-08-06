package demo_ucloud

import "fmt"

func describeElb() (string, string) {
	elbCli := newElbCli()
	describeULBRequest := elbCli.NewDescribeULBRequest()
	describeULBRequest.Region = String(region)
	response, err := elbCli.DescribeULB(describeULBRequest)
	if err != nil {
		panic(err)
	}
	if response.TotalCount != 1 {
		panic("elb count not one")
	}
	ulbInstance := response.DataSet[0]
	fmt.Println(ulbInstance)
	ulbvServerSet := ulbInstance.VServerSet[0]
	return ulbInstance.ULBId, ulbvServerSet.VServerId
}

func cleanElb(array []string) {
	elb, _ := describeElb()
	elbCli := newElbCli()
	for _, host := range array {
		request := elbCli.NewReleaseBackendRequest()
		request.Region = String(region)
		request.ULBId = String(elb)
		request.BackendId = String(host)
		backend, err := elbCli.ReleaseBackend(request)
		if err != nil {
			panic(err)
		}
		fmt.Println(backend)
	}
}

func changeUlbServer(array []string) {
	elb, vServerId := describeElb()
	elbCli := newElbCli()
	for _, host := range array {
		request := elbCli.NewAllocateBackendRequest()
		request.Region = String(region)
		request.ULBId = &elb
		request.VServerId = &vServerId
		request.ResourceType = String("UHost")
		request.ResourceId = &host
		request.Port = Int(6443)
		backend, err := elbCli.AllocateBackend(request)
		if err != nil {
			panic(err)
		}
		fmt.Println(backend)
	}
}
