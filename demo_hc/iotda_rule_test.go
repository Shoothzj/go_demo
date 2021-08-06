package demo_hc

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	iotda "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotda/v5"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotda/v5/model"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotda/v5/region"
	"go_demo/component_hc_config"
	"testing"
)

func TestCreateRule(t *testing.T) {
	client := getIotda()
	actions := make([]model.RuleAction, 1)
	actions[0] = model.RuleAction{
		Type:            "",
		Addition:        nil,
		SmnForwarding:   nil,
		DeviceAlarm:     nil,
		DeviceCommand:   nil,
		DisForwarding:   nil,
		ObsForwarding:   nil,
		RomaForwarding:  nil,
		IotaForwarding:  nil,
		KafkaForwarding: nil,
	}
	rule := &model.Rule{
		Name:           "obs-test",
		Description:    nil,
		ConditionGroup: nil,
		Actions:        actions,
		RuleType:       "active",
		Status:         nil,
		AppId:          nil,
		EdgeNodeIds:    nil,
	}
	req := &model.CreateRuleRequest{Body: rule}
	resp, err := client.CreateRule(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func TestQueryAllRules(t *testing.T) {
	hcConfig := component_hc_config.ReadHcConfig()
	auth := basic.NewCredentialsBuilder().WithAk(hcConfig.Ak).WithSk(hcConfig.Sk).Build()
	client := iotda.NewIoTDAClient(iotda.IoTDAClientBuilder().WithRegion(region.CN_NORTH_4).WithCredential(auth).Build())
	var limit int32 = 10
	request := &model.ListRulesRequest{Limit: &limit}
	listRules, err := client.ListRules(request)
	if err != nil {
		panic(err)
	}
	rules := *listRules.Rules
	fmt.Println("length is ", len(rules))
	for _, rule := range rules {
		fmt.Println(rule)
	}
}
