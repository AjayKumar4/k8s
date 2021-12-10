package utils

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// UseCurrentContextConfig - return current context of kubeconfig
func UseCurrentContextConfig(configFilepath string) *rest.Config {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", configFilepath)
	if err != nil {
		panic(err.Error())
	}
	return config
}

// CreateK8sClientSet - return kubernetes client set
func CreateK8sClientSet(configFilepath string) *kubernetes.Clientset {
	config := UseCurrentContextConfig(configFilepath)
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}
