package main

import (
	"fmt"
	"github.com/zouyx/agollo/v4"
	"github.com/zouyx/agollo/v4/env/config"
	"testing"
)

func TestGetConfigFromApolloHttp(t *testing.T) {
	c := &config.AppConfig{
		AppID:         appId,
		Cluster:       cluster,
		IP:            apolloConnStr,
		NamespaceName: namespace,
	}
	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	if err != nil {
		panic(err)
	}
	intValue := client.GetConfigAndInit(namespace).GetIntValue("timeout", 0)
	fmt.Println(intValue)
}

func TestGetConfigFromApolloHttps(t *testing.T) {
	c := &config.AppConfig{
		AppID:         appId,
		Cluster:       cluster,
		IP:            apolloConnStrHttps,
		NamespaceName: namespace,
	}
	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	if err != nil {
		panic(err)
	}
	value := client.GetConfigAndInit(namespace).GetValue("timeout")
	fmt.Println(value)
}
