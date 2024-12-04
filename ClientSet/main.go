package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 获取deployment
	// 1、先获取configconfig
	config, err := clientcmd.BuildConfigFromFlags("", "../config/kubeconfig.conf")
	if err != nil {
		panic(err)
	}

	//2、创建client client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	deployment, _ := clientset.AppsV1().Deployments("krm").Get(context.TODO(), "krm-backend", metav1.GetOptions{})
	fmt.Printf("当前deploymeng的镜像是: %s \n", deployment.Spec.Template.Spec.Containers[0].Image)
}
