package util

import "k8s.io/client-go/kubernetes"

type Manager struct {
	kubeclient *kubernetes.Clientset
}

func NewManager(kubeclient *kubernetes.Clientset) *Manager {
	return &Manager{
		kubeclient: kubeclient,
	}
}
