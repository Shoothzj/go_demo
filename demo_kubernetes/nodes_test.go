package demo_kubernetes

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"testing"
)

func TestShowNodes(t *testing.T) {
	// default path
	var path = filepath.Join(homedir.HomeDir(), ".kube", "config")
	fmt.Println(path)
	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	list, err := clientset.CoreV1().Nodes().List(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(list)
}
