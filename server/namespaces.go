package server

import (
	"encoding/json"
	"net/http"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog"
)

func (s *Server) getNamespaces(w http.ResponseWriter, r *http.Request) {
	namespaces, err := s.manager.GetNamespaces()
	if err != nil {
		klog.Errorf("Error while getting the namespaces: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var allowedNamespaces []corev1.Namespace

	// filter namespaces
	for _, namespace := range namespaces {
		if isAllowed(namespace.Name, s.namespaceRegexes) {
			allowedNamespaces = append(allowedNamespaces, namespace)
		}
	}

	err = json.NewEncoder(w).Encode(allowedNamespaces)

	if err != nil {
		klog.Errorf("Error while encoding the namespaces: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
