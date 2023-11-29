package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mheers/k8s-secret-ui/pkg/util"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog"
)

func (s *Server) createSecret(w http.ResponseWriter, r *http.Request) {
	var secret corev1.Secret
	err := json.NewDecoder(r.Body).Decode(&secret)
	if err != nil {
		klog.Errorf("Error while decoding the secret: %s", err.Error())
		w.Write([]byte(err.Error()))
		return
	}

	// check if secret is allowed
	if !isAllowed(secret.Name, s.secretRegexes) || !isAllowed(secret.Namespace, s.namespaceRegexes) {
		klog.Errorf("Secret %s is not allowed in namespace %s", secret.Name, secret.Namespace)
		w.Write([]byte("Secret is not allowed"))
		return
	}

	res := util.CreateSecret(kubeclient, secret)

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		klog.Errorf("Error while encoding the secret: %s", err.Error())
		w.Write([]byte(err.Error()))
		return
	}
}

func (s *Server) updateSecret(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	secretName := pathParams["secretname"]
	secretNamespace := pathParams["secretns"]

	var secret corev1.Secret
	err := json.NewDecoder(r.Body).Decode(&secret)
	if err != nil {
		klog.Errorf("Error while decoding the secret: %s", err.Error())
		w.Write([]byte(err.Error()))
		return
	}

	// check if secret name and namespace match the path params
	if secretName != secret.Name || secretNamespace != secret.Namespace {
		klog.Errorf("Secret name or namespace does not match the path params")
		w.Write([]byte("Secret name or namespace does not match the path params"))
		return
	}

	// check if secret is allowed
	if !isAllowed(secret.Name, s.secretRegexes) || !isAllowed(secret.Namespace, s.namespaceRegexes) {
		klog.Errorf("Secret %s is not allowed in namespace %s", secret.Name, secret.Namespace)
		w.Write([]byte("Secret is not allowed"))
		return
	}

	err = json.NewEncoder(w).Encode(util.UpdateSecret(kubeclient, secret))
	if err != nil {
		klog.Errorf("Error while encoding the secret: %s", err.Error())
		w.Write([]byte(err.Error()))
		return
	}
}

func (s *Server) getSecret(w http.ResponseWriter, r *http.Request) {
	queryParams := mux.Vars(r)
	secretName := queryParams["secretname"]
	secretNamespace := queryParams["secretns"]

	// check if secret is allowed
	if !isAllowed(secretName, s.secretRegexes) || !isAllowed(secretNamespace, s.namespaceRegexes) {
		klog.Errorf("Secret %s is not allowed in namespace %s", secretName, secretNamespace)
		w.Write([]byte("Secret is not allowed"))
		return
	}

	err := json.NewEncoder(w).Encode(util.GetSecret(kubeclient, secretNamespace, secretName))
	if err != nil {
		klog.Errorf("Error while encoding the secret data: %s", err.Error())
		w.Write([]byte(err.Error()))
		return
	}
}

func (s *Server) getSecretsOfNS(w http.ResponseWriter, r *http.Request) {
	queryParams := mux.Vars(r)
	namespace := queryParams["secretns"]

	// check if namespace is allowed
	if !isAllowed(namespace, s.namespaceRegexes) {
		klog.Errorf("Namespace %s is not allowed", namespace)
		w.Write([]byte("Namespace is not allowed"))
		return
	}

	err := json.NewEncoder(w).Encode(util.GetSecretsOfNS(kubeclient, namespace))
	if err != nil {
		klog.Errorf("Error while encoding the secrets: %s", err.Error())
		w.Write([]byte(err.Error()))
		return
	}
}

func (s *Server) deleteSecret(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	secretName := pathParams["secretname"]
	secretNamespace := pathParams["secretns"]

	// check if secret is allowed
	if !isAllowed(secretName, s.secretRegexes) || !isAllowed(secretNamespace, s.namespaceRegexes) {
		klog.Errorf("Secret %s is not allowed in namespace %s", secretName, secretNamespace)
		w.Write([]byte("Secret is not allowed"))
		return
	}

	res := util.DeleteSecret(kubeclient, secretNamespace, secretName)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		klog.Errorf("Error while encoding the secret: %s", err.Error())
		w.Write([]byte(err.Error()))
		return
	}
}
