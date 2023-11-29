package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog"
)

func (s *Server) createSecret(w http.ResponseWriter, r *http.Request) {
	var secret corev1.Secret
	err := json.NewDecoder(r.Body).Decode(&secret)
	if err != nil {
		klog.Errorf("Error while decoding the secret: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check if secret is allowed
	if !isAllowed(secret.Name, s.secretRegexes) || !isAllowed(secret.Namespace, s.namespaceRegexes) {
		klog.Errorf("Secret %s is not allowed in namespace %s", secret.Name, secret.Namespace)
		http.Error(w, "Secret is not allowed", http.StatusForbidden)
		return
	}

	secretCreated, err := s.manager.CreateSecret(secret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(secretCreated)
	if err != nil {
		klog.Errorf("Error while encoding the secret: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (s *Server) updateSecret(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	secretName := pathParams["secretname"]
	secretNamespace := pathParams["secretns"]

	var secret *corev1.Secret
	err := json.NewDecoder(r.Body).Decode(secret)
	if err != nil {
		klog.Errorf("Error while decoding the secret: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check if secret name and namespace match the path params
	if secretName != secret.Name || secretNamespace != secret.Namespace {
		klog.Errorf("Secret name or namespace does not match the path params")
		http.Error(w, "Secret name or namespace does not match the path params", http.StatusBadRequest)
		return
	}

	// check if secret is allowed
	if !isAllowed(secret.Name, s.secretRegexes) || !isAllowed(secret.Namespace, s.namespaceRegexes) {
		klog.Errorf("Secret %s is not allowed in namespace %s", secret.Name, secret.Namespace)
		http.Error(w, "Secret is not allowed", http.StatusForbidden)
		return
	}

	secretUpdated, err := s.manager.UpdateSecret(secret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(secretUpdated)
	if err != nil {
		klog.Errorf("Error while encoding the secret: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
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
		http.Error(w, "Secret is not allowed", http.StatusForbidden)
		return
	}

	secret, err := s.manager.GetSecret(secretNamespace, secretName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(secret)
	if err != nil {
		klog.Errorf("Error while encoding the secret data: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (s *Server) getSecretsOfNS(w http.ResponseWriter, r *http.Request) {
	queryParams := mux.Vars(r)
	namespace := queryParams["secretns"]

	// check if namespace is allowed
	if !isAllowed(namespace, s.namespaceRegexes) {
		klog.Errorf("Namespace %s is not allowed", namespace)
		http.Error(w, "Namespace is not allowed", http.StatusForbidden)
		return
	}

	secrets, err := s.manager.GetSecretsOfNS(namespace)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	allowedSecrets := []corev1.Secret{}
	for _, secret := range secrets {
		if isAllowed(secret.Name, s.secretRegexes) {
			allowedSecrets = append(allowedSecrets, secret)
		}
	}

	err = json.NewEncoder(w).Encode(allowedSecrets)
	if err != nil {
		klog.Errorf("Error while encoding the secrets: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
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
		http.Error(w, "Secret is not allowed", http.StatusForbidden)
		return
	}

	err := s.manager.DeleteSecret(secretNamespace, secretName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Secret deleted"))
}
