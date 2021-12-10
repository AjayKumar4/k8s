package main

import (
	"context"
	"fmt"
	"os"

	"k8s/utils"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	versionNo string // version number of program (from VERSION file)
	buildTime string // when the executable was built
	buildUser string // user who built the program
)

func main() {

	opt := utils.Init()

	// print version number and exit
	if opt.GetVersion() {
		fmt.Printf("version %s \nbuild on %s \nbuild by %s", versionNo, buildTime, buildUser)
		os.Exit(0)
	}

	// create the clientset
	clientset := utils.CreateK8sClientSet(opt.GetKubeConfig())

	deploymentsClient := clientset.AppsV1().Deployments(opt.GetNamespace())

	fmt.Printf("Listing deployments in namespace %q:\n", opt.GetNamespace())
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
