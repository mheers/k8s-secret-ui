package util

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
)

func GetSecretsOfNS(kubeclient *kubernetes.Clientset, namespace string) []corev1.Secret {
	secrets, err := kubeclient.CoreV1().Secrets(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		klog.Errorf("Error getting the secrets from a namespace %s", err.Error())
	}

	return secrets.Items
}

func GetSecretData(kubeclient *kubernetes.Clientset, secreteNS, secretName string) corev1.Secret {
	secret, err := kubeclient.CoreV1().Secrets(secreteNS).Get(context.Background(), secretName, metav1.GetOptions{})
	if err != nil {
		klog.Errorf("Error getting secret data %s", err.Error())
	}

	return *secret
}

func UpdateSecret(kubeclient *kubernetes.Clientset, secretNS, secretName string, secret corev1.Secret) *corev1.Secret {
	sec, err := kubeclient.CoreV1().Secrets(secretNS).Update(context.Background(), &secret, metav1.UpdateOptions{})
	if err != nil {
		klog.Errorf("Error updating the secret %s", err.Error())
	}

	return sec
}

func DeleteSecret(kubeclient *kubernetes.Clientset, secretNS, secretName string) error {
	return kubeclient.CoreV1().Secrets(secretNS).Delete(context.Background(), secretName, metav1.DeleteOptions{})
}

func CreateSecret(kubeclient *kubernetes.Clientset, secret corev1.Secret) error {
	_, err := kubeclient.CoreV1().Secrets(secret.Namespace).Create(context.Background(), &secret, metav1.CreateOptions{})
	return err
}
