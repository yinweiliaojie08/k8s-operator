package main

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 获取deployment
	// RESTClient创建
	// 1、先获取configconfig
	config, err := clientcmd.BuildConfigFromFlags("", "../config/kubeconfig.conf")
	if err != nil {
		panic(err)
	}
	config.GroupVersion = &appsv1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	config.APIPath = "/apis"

	//2、创建client client
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	// get deployment
	deployment := appsv1.Deployment{}
	err = restClient.Get().Namespace("krm").Resource("deployments").Name("krm-backend").Do(context.TODO()).Into(&deployment)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(deployment.Kind)
		fmt.Println(deployment.Spec.Selector)
		fmt.Println(deployment.Namespace)
	}
	// ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	// 获取pod
	// 1、先获取configconfig
	config2, err := clientcmd.BuildConfigFromFlags("", "../config/kubeconfig.conf")
	if err != nil {
		panic(err)
	}
	config2.GroupVersion = &corev1.SchemeGroupVersion
	config2.NegotiatedSerializer = scheme.Codecs
	config2.APIPath = "/api"

	//2、创建client client
	restClient2, err := rest.RESTClientFor(config2)
	if err != nil {
		panic(err)
	}

	// get pod
	pod := corev1.Pod{}
	err = restClient2.Get().Namespace("krm").Resource("pods").Name("krm-backend-7d7b5fcd86-kjc4g").Do(context.TODO()).Into(&pod)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pod.Kind)
		fmt.Println(pod.Name)
		fmt.Println(pod.Namespace)
	}

}
