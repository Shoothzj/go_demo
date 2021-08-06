package demo_etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

const timeout = time.Second * 4

func TestEtcdClient(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer cli.Close()
}

func TestGet(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	response, err := cli.Get(ctx, "sample_key")
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
	cancel()
}

func TestGetPut(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	defer cli.Close()
	if err != nil {
		panic(err)
	}
	timeout := time.Second * 4
	{
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		response, err := cli.Put(ctx, "sample_key1", "sample_value")
		cancel()
		if err != nil {
			panic(err)
		}
		fmt.Println(response)
	}
	{
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		response, err := cli.Get(ctx, "sample_key1")
		cancel()
		if err != nil {
			panic(err)
		}
		header := response.Header
		fmt.Println("raft revision is ", header.Revision)
		keyValue := response.Kvs[0]
		fmt.Println("data version is ", keyValue.Version)
		fmt.Println(response)
	}
}

func TestTransaction(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	defer cli.Close()
	if err != nil {
		panic(err)
	}
	timeout := time.Second * 4
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	cli.KV.Txn(ctx).If(clientv3.Compare(clientv3.Value("sample_key"), "=", "sample_value")).Then(clientv3.OpPut("sample_key", "2")).Commit()
	cancel()
}

// TestWatch
func TestWatch(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	// 使用 WithCancel 构造一个带控制的 context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rch := cli.Watch(ctx, "", clientv3.WithPrefix(), clientv3.WithPrevKV())
	for n := range rch {
		//rch is a remote channel
		for _, ev := range n.Events {
			switch ev.Type {
			case mvccpb.PUT:
				//do something with event ADD
			case mvccpb.DELETE:
				//do something with event DEL
				fmt.Println("find DETELE:", ev.PrevKv.Key, ev.PrevKv.Value)
			}
		}
	}
}
