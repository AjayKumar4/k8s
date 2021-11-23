package main

import (
	"context"
	"fmt"

	"kubernetes-api-go/utils"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func main() {

	opt := utils.Init()

	// create the clientset
	clientset := utils.CreateK8sClientSet(opt)

	deploymentsClient := clientset.AppsV1().Deployments(*opt.GetNamespace())

	fmt.Printf("Listing deployments in namespace %q:\n", *opt.GetNamespace())
	list, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range list.Items {
		containers := d.Spec.Template.Spec.Containers
		for _, container := range containers {
			fmt.Printf(" * ( deployment %s ) (replicas %d ) (image %s) (creation time %s)\n", d.Name, *d.Spec.Replicas, container.Image, d.CreationTimestamp.String())
		}
	}
}
