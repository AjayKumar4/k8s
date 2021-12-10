package utils

import (
	"flag"
	"path/filepath"

	"k8s.io/client-go/util/homedir"
)

// Options - Configuration taken from the argument
type Options struct {
	kubeconfig *string
	namespace  *string
	version    *bool
}

// GetKubeConfig - return kubeconfig from the argument
func (o *Options) GetKubeConfig() string {
	return *o.kubeconfig
}

// GetNamespace - return namespace from the argument
func (o *Options) GetNamespace() string {
	return *o.namespace
}

// GetVersion - return version from the argument
func (o *Options) GetVersion() bool {
	return *o.version
}

// Init - retrieve options from the argument
func Init() *Options {
	opt := new(Options)

	if home := homedir.HomeDir(); home != "" {
		opt.kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		opt.kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	opt.namespace = flag.String("namespace", "", "(optional) default namespace")
	opt.version = flag.Bool("version", false, "print out the version number and build date of the program")
	flag.Parse()

	return opt
}
