package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mheers/k8s-secret-ui/pkg/util"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog"
)

func (s *Server) createConfigMap(w http.ResponseWriter, r *http.Request) {
	var configMap corev1.ConfigMap
	err := json.NewDecoder(r.Body).Decode(&configMap)
	if err != nil {
		klog.Errorf("Error while decoding the configmap: %s", err.Error())
		w.Write([]byte(err.Error()))
		return
	}

	// check if configMap is allowed
	if !isAllowed(configMap.Name, s.configMapRegexes) || !isAllowed(configMap.Namespace, s.namespaceRegexes) {
		klog.Errorf("ConfigMap %s is not allowed in namespace %s", configMap.Name, configMap.Namespace)
		w.Write([]byte("ConfigMap is not allowed"))
		return
	}

	res := util.CreateConfigMap(kubeclient, configMap)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		klog.Errorf("Error while encoding the configmap: %s", err.Error())
		w.Write([]byte(err.Error()))
		return
	}
}

func (s *Server) getConfigMapsOfNS(w http.ResponseWriter, r *http.Request) {
	queryParams := mux.Vars(r)
	namespace := queryParams["cmns"]

	// check if namespace is allowed
	if !isAllowed(namespace, s.namespaceRegexes) {
		klog.Errorf("Namespace %s is not allowed", namespace)
		w.Write([]byte("Namespace is not allowed"))
		return
	}

	err := json.NewEncoder(w).Encode(util.GetConfigMapsOfNS(kubeclient, namespace))
	if err != nil {
		klog.Errorf("Error while encoding the configmaps: %s", err.Error())
		w.Write([]byte(err.Error()))
		return
	}
}

func (s *Server) getConfigMap(w http.ResponseWriter, r *http.Request) {
	queryParams := mux.Vars(r)
	configMapName := queryParams["cmname"]
	configMapNamespace := queryParams["cmns"]

	// check if configMap is allowed
	if !isAllowed(configMapName, s.configMapRegexes) || !isAllowed(configMapNamespace, s.namespaceRegexes) {
		klog.Errorf("ConfigMap %s is not allowed in namespace %s", configMapName, configMapNamespace)
		w.Write([]byte("ConfigMap is not allowed"))
		return
	}

	err := json.NewEncoder(w).Encode(util.GetConfigMap(kubeclient, configMapNamespace, configMapName))
	if err != nil {
		klog.Errorf("Error while encoding the configmap: %s", err.Error())
		w.Write([]byte(err.Error()))
		return
	}
}

func (s *Server) updateConfigMap(w http.ResponseWriter, r *http.Request) {
	queryParams := mux.Vars(r)
	configMapName := queryParams["cmname"]
	configMapNamespace := queryParams["cmns"]

	var configMap corev1.ConfigMap
	err := json.NewDecoder(r.Body).Decode(&configMap)
	if err != nil {
		klog.Errorf("Error while decoding the configmap: %s", err.Error())
		w.Write([]byte(err.Error()))
		return
	}

	// check if configMap name and namespace match the path params
	if configMapName != configMap.Name || configMapNamespace != configMap.Namespace {
		klog.Errorf("ConfigMap name or namespace does not match the path params")
		w.Write([]byte("ConfigMap name or namespace does not match the path params"))
		return
	}

	// check if configMap is allowed
	if !isAllowed(configMap.Name, s.secretRegexes) || !isAllowed(configMap.Namespace, s.namespaceRegexes) {
		klog.Errorf("ConfigMap %s is not allowed in namespace %s", configMap.Name, configMap.Namespace)
		w.Write([]byte("ConfigMap is not allowed"))
		return
	}

	err = json.NewEncoder(w).Encode(util.UpdateConfigMap(kubeclient, configMapNamespace, configMapName, configMap))
	if err != nil {
		klog.Errorf("Error while encoding the configmap: %s", err.Error())
		w.Write([]byte(err.Error()))
		return
	}

}

func (s *Server) deleteConfigMap(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	configMapName := pathParams["cmname"]
	configMapNamespace := pathParams["cmns"]

	// check if configMap is allowed
	if !isAllowed(configMapName, s.configMapRegexes) || !isAllowed(configMapNamespace, s.namespaceRegexes) {
		klog.Errorf("ConfigMap %s is not allowed in namespace %s", configMapName, configMapNamespace)
		w.Write([]byte("ConfigMap is not allowed"))
		return
	}

	res := util.DeleteConfigMap(kubeclient, configMapNamespace, configMapName)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		klog.Errorf("Error while encoding the configmap: %s", err.Error())
		w.Write([]byte(err.Error()))
		return
	}
}
