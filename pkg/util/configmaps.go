package util

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
)

func (m *Manager) ListConfigMaps() ([]corev1.ConfigMap, error) {
	configMapList, err := m.kubeclient.CoreV1().ConfigMaps("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		klog.Errorf("Error while listing all the configmaps: %s", err.Error())
		return nil, err
	}

	return configMapList.Items, nil
}

func (m *Manager) GetConfigMap(cmns, cmName string) (*corev1.ConfigMap, error) {
	cm, err := m.kubeclient.CoreV1().ConfigMaps(cmns).Get(context.Background(), cmName, metav1.GetOptions{})
	if err != nil {
		klog.Errorf("Error while getting a configmap: %s", err.Error())
		return nil, err
	}
	return cm, nil
}

func (m *Manager) UpdateConfigMap(configmap *corev1.ConfigMap) (*corev1.ConfigMap, error) {
	cm, err := m.kubeclient.CoreV1().ConfigMaps(configmap.Namespace).Update(context.Background(), configmap, metav1.UpdateOptions{})
	if err != nil {
		klog.Errorf("Error while updating the configmap: %s", err.Error())
		return nil, err
	}
	return cm, nil
}

func (m *Manager) GetNamespaces() ([]corev1.Namespace, error) {
	allNamespaces, err := m.kubeclient.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		klog.Errorf("Error while getting all the NSs %s", err.Error())
		return nil, err
	}
	return allNamespaces.Items, nil
}

func (m *Manager) GetConfigMapsOfNS(namespace string) ([]corev1.ConfigMap, error) {
	configMaps, err := m.kubeclient.CoreV1().ConfigMaps(namespace).List(context.Background(), metav1.ListOptions{
		Limit: 100,
	})
	if err != nil {
		klog.Errorf("Error while listing all CMs of a NS %s", err.Error())
		return nil, err
	}
	return configMaps.Items, nil
}

func (m *Manager) DeleteConfigMap(namespace, name string) error {
	return m.kubeclient.CoreV1().ConfigMaps(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
}

func (m *Manager) CreateConfigMap(cm corev1.ConfigMap) (*corev1.ConfigMap, error) {
	return m.kubeclient.CoreV1().ConfigMaps(cm.Namespace).Create(context.Background(), &cm, metav1.CreateOptions{})
}
