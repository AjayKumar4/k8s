package utils

import (
	"flag"
	"path/filepath"

	"k8s.io/client-go/util/homedir"
)

type Options struct {
	kubeconfig *string
	namespace  *string
}

func (o *Options) GetKubeConfig() *string {

	return o.kubeconfig
}

func (o *Options) GetNamespace() *string {

	flag.Parse()
	return o.namespace
}

func Init() *Options {
	opt := new(Options)

	if home := homedir.HomeDir(); home != "" {
		opt.kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		opt.kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	opt.namespace = flag.String("namespace", "", "(optional) default namespace")
	flag.Parse()

	return opt
}
