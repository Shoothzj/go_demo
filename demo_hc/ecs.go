package demo_hc

import (
	"fmt"
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
)

func getEcsClient() *ecs.EcsClient {
	return ecs.NewEcsClient(
		ecs.EcsClientBuilder().
			WithRegion(region.ValueOf("ap-southeast-1")).
			WithCredential(*getAuth()).
			Build())
}

func CreateEcs() {
	client := ecs.NewEcsClient(
		ecs.EcsClientBuilder().
			WithRegion(region.ValueOf("ap-southeast-1")).
			WithCredential(*getAuth()).
			Build())
	request := &model.CreateServersRequest{}
	sizeRootVolumePrePaidServerRootVolume:= int32(40)
	rootVolumeServer := &model.PrePaidServerRootVolume{
		Volumetype: model.GetPrePaidServerRootVolumeVolumetypeEnum().SATA,
		Size: &sizeRootVolumePrePaidServerRootVolume,
	}
	chargingModeExtendparamPrePaidServerEipExtendParam:= model.GetPrePaidServerEipExtendParamChargingModeEnum().POST_PAID
	extendparamEip := &model.PrePaidServerEipExtendParam{
		ChargingMode: &chargingModeExtendparamPrePaidServerEipExtendParam,
	}
	sizeBandwidthPrePaidServerEipBandwidth:= int32(1)
	bandwidthEip := &model.PrePaidServerEipBandwidth{
		Size: &sizeBandwidthPrePaidServerEipBandwidth,
	}
	eipPublicip := &model.PrePaidServerEip{
		Iptype: "5_bgp",
		Bandwidth: bandwidthEip,
		Extendparam: extendparamEip,
	}
	publicipServer := &model.PrePaidServerPublicip{
		Eip: eipPublicip,
	}
	var listNicsServer = []model.PrePaidServerNic{
		{
			SubnetId: "subnetid",
		},
	}
	serverbody := &model.PrePaidServer{
		ImageRef: "67159f00-89c1-44eb-b810-c90478376cb4",
		FlavorRef: "c6.large.2",
		Name: "network",
		Vpcid: "vpcid",
		Nics: listNicsServer,
		Publicip: publicipServer,
		RootVolume: rootVolumeServer,
		AvailabilityZone: "az",
	}
	dryRunCreateServersRequestBody:= false
	request.Body = &model.CreateServersRequestBody{
		Server: serverbody,
		DryRun: &dryRunCreateServersRequestBody,
	}
	response, err := client.CreateServers(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}

func QueryEcs() *Ecs {
	client := ecs.NewEcsClient(
		ecs.EcsClientBuilder().
			WithRegion(region.ValueOf("ap-southeast-1")).
			WithCredential(*getAuth()).
			Build())

	request := &model.ListServersDetailsRequest{}
	response, err := client.ListServersDetails(request)
	if err != nil {
		panic(err)
	}
	listServersDetailsResponse := response
	if *listServersDetailsResponse.Count != 1 {
		panic(fmt.Sprintf("server count is %d", listServersDetailsResponse.Count))
	}
	serverDetail := (*listServersDetailsResponse.Servers)[0]
	fmt.Println(serverDetail)
	e := &Ecs{Id: serverDetail.Id}
	return e;
}

func DelEcs() {
	ecsClient := getEcsClient()
	queryEcs := QueryEcs()
	request := &model.DeleteServersRequest{}
	var listServersbody = []model.ServerId{
		{
			Id: queryEcs.Id,
		},
	}
	request.Body = &model.DeleteServersRequestBody{
		Servers: listServersbody,
	}
	response, err := ecsClient.DeleteServers(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}

