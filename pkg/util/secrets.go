package util

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

func (m *Manager) GetSecretsOfNS(namespace string) ([]corev1.Secret, error) {
	secrets, err := m.kubeclient.CoreV1().Secrets(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		klog.Errorf("Error getting the secrets from a namespace %s", err.Error())
		return nil, err
	}

	return secrets.Items, nil
}

func (m *Manager) GetSecret(secreteNS, secretName string) (*corev1.Secret, error) {
	secret, err := m.kubeclient.CoreV1().Secrets(secreteNS).Get(context.Background(), secretName, metav1.GetOptions{})
	if err != nil {
		klog.Errorf("Error getting secret data %s", err.Error())
		return nil, err
	}

	return secret, nil
}

func (m *Manager) UpdateSecret(secret *corev1.Secret) (*corev1.Secret, error) {
	sec, err := m.kubeclient.CoreV1().Secrets(secret.Namespace).Update(context.Background(), secret, metav1.UpdateOptions{})
	if err != nil {
		klog.Errorf("Error updating the secret %s", err.Error())
		return nil, err
	}

	return sec, nil
}

func (m *Manager) DeleteSecret(secretNS, secretName string) error {
	return m.kubeclient.CoreV1().Secrets(secretNS).Delete(context.Background(), secretName, metav1.DeleteOptions{})
}

func (m *Manager) CreateSecret(secret corev1.Secret) (*corev1.Secret, error) {
	return m.kubeclient.CoreV1().Secrets(secret.Namespace).Create(context.Background(), &secret, metav1.CreateOptions{})
}
