package k8s

import (
	"errors"
	"fmt"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
)

func GetK8sClient() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		// If running outside of the cluster, try to load the configuration from
		// the KUBECONFIG environment variable.
		kubeconfigPath := os.Getenv("KUBECONFIG")
		if kubeconfigPath == "" {
			klog.Errorf("could not load Kubernetes configuration from env")
		}

		config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		if err != nil {
			klog.Errorf("could not load Kubernetes configuration from %s: %v", kubeconfigPath, err)

			// check if $HOME/.kube/config exists
			if _, err := os.Stat(fmt.Sprintf("%s/.kube/config", os.Getenv("HOME"))); err == nil {
				config, err = clientcmd.BuildConfigFromFlags("", fmt.Sprintf("%s/.kube/config", os.Getenv("HOME")))
				if err != nil {
					klog.Errorf("could not load Kubernetes configuration from %s: %v", fmt.Sprintf("%s/.kube/config", os.Getenv("HOME")), err)
				}
			}
		}

		if config == nil {
			return nil, errors.New("could not load Kubernetes configuration")
		}
	}

	// Create a Kubernetes client.
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("could not create Kubernetes client: %v", err)
	}

	return clientset, nil
}
