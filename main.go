package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/home/william/.kube/config", "kubeconfig file location")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatalf("error %s building config from flags\n", err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("error %s creating the clientset\n", err.Error())
	}

	ctx := context.Background()
	pods, err := clientset.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Fatalf("error %s while listing all the pods from default namespace\n", err.Error())
	}

	for _, pod := range pods.Items {
		fmt.Printf("%s\n", pod.Name)
	}

	deployments, err := clientset.AppsV1().Deployments("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Fatalf("error %s listing the deployments\n", err.Error())
	}

	for _, d := range deployments.Items {
		fmt.Printf("%s\n", d.Name)
	}
}
