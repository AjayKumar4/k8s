package utils

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func UseCurrentContextConfig(options *Options) *rest.Config {
	kubeconfig := options.GetKubeConfig()
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	return config
}

func CreateK8sClientSet(options *Options) *kubernetes.Clientset {
	config := UseCurrentContextConfig(options)
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}
