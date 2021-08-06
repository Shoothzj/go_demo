package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var deletePeriod = flag.Int("deletePeriod", 30, "Delete periods in seconds")
var deletePodList = flag.String("deletePodList", "", "Input your wanted delete pod, split it in ,")

func main() {
	flag.Parse()
	homeDir := homedir.HomeDir()
	kubeConfigPath := homeDir + "/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	ticker := time.NewTicker(time.Duration(*deletePeriod) * time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
			pods, err := clientset.CoreV1().Pods("default").List(metav1.ListOptions{})
			if err != nil {
				panic(err.Error())
			}
			split := strings.Split(*deletePodList, ",")
			for _, pod := range pods.Items {
				podName := pod.GetName()
				for i := range split {
					if strings.Contains(podName, split[i]) {
						helper := int64(0)
						deleteOptions := metav1.DeleteOptions{GracePeriodSeconds: &helper}
						fmt.Println("Deleted ", podName)
						clientset.CoreV1().Pods("default").Delete(podName, &deleteOptions)
					}
				}
			}
		}
	}

}
