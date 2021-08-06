package demo_ucloud

import (
	"fmt"
	"testing"
)

func TestCreateEcs(t *testing.T) {
	uHostClient := newHostCli()
	for i := 0; i < 1; i++ {
		createUHostInstanceRequest := uHostClient.NewCreateUHostInstanceRequest()
		createUHostInstanceRequest.Region = String(region)
		createUHostInstanceRequest.ImageId = String("bsi-pog1ebfx")
		createUHostInstanceRequest.CPU = Int(4)
		createUHostInstanceRequest.ChargeType = String("Dynamic")
		createUHostInstanceRequest.Memory = Int(16384)
		createUHostInstanceRequest.Zone = String("hk-02")
		createUHostInstanceRequest.MachineType = String("N4")
		hostInstance, err := uHostClient.CreateUHostInstance(createUHostInstanceRequest)
		if err != nil {
			panic(err)
		}
		fmt.Println(hostInstance)
	}
}

func TestDescribeEcs(t *testing.T) {
	hostSet := listEcs()
	for _, host := range hostSet {
		fmt.Println(host)
	}
	fmt.Println(hostSet)
}

func TestShutdownAllEcs(t *testing.T) {
	uHostClient := newHostCli()
	hostSet := listEcs()
	for _, host := range hostSet {
		stopUHostInstanceRequest := uHostClient.NewStopUHostInstanceRequest()
		stopUHostInstanceRequest.Region = String(region)
		stopUHostInstanceRequest.UHostId = String(host.UHostId)
		terminateUHostInstance, err := uHostClient.StopUHostInstance(stopUHostInstanceRequest)
		if err != nil {
			panic(err)
		}
		fmt.Println(terminateUHostInstance)
	}
}

func TestDeleteAllEcs(t *testing.T) {
	uHostClient := newHostCli()
	hostSet := listEcs()
	for _, host := range hostSet {
		terminateUHostInstanceRequest := uHostClient.NewTerminateUHostInstanceRequest()
		terminateUHostInstanceRequest.Region = String(region)
		terminateUHostInstanceRequest.UHostId = String(host.UHostId)
		terminateUHostInstance, err := uHostClient.TerminateUHostInstance(terminateUHostInstanceRequest)
		if err != nil {
			panic(err)
		}
		fmt.Println(terminateUHostInstance)
	}
}
